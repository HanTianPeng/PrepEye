import java.io.IOException;
import java.nio.file.*;
import java.util.*;
import java.util.regex.*;
import java.util.stream.*;

/**
 * <h1>最小化 Skill 引擎 —— 纯 Java 零依赖 Demo</h1>
 *
 * <p>演示 Skill 系统的 5 个核心环节：</p>
 * <pre>
 * ┌──────────────────────────────────────────────────────────────────┐
 * │  环节 1: 加载 —— 扫描 skills/ 目录，解析 SKILL.md               │
 * │  环节 2: 匹配 —— 用户输入 → 关键词匹配 → 激活哪些 Skill          │
 * │  环节 3: 注入 —— 把 Skill body 拼到 System Prompt 末尾           │
 * │  环节 4: 限制 —— 应用 allowed-tools 权限约束                     │
 * │  环节 5: 演示 —— 对比"有 Skill"和"没有 Skill"的 prompt 差异       │
 * └──────────────────────────────────────────────────────────────────┘
 * </pre>
 *
 * <h2>运行方式</h2>
 * <pre>java SkillEngineDemo.java</pre>
 *
 * <h2>核心设计原则（从 Claude Code 迁移）</h2>
 * <pre>
 *   1. YAML Frontmatter 给 Runtime 读（机器的配置）
 *   2. Markdown Body 给 LLM 读（模型的指令）
 *   3. description 是匹配的唯一依据（模型用它判断"该不该激活"）
 *   4. Skill 注入不破坏 prompt cache（追加在 system prompt 尾部）
 * </pre>
 *
 * <h2>Skill 的物理形态</h2>
 * <pre>
 *   skills/
 *   ├── code-review/
 *   │   └── SKILL.md          ← 一个 Skill = 一个目录 + 一个 SKILL.md
 *   ├── deploy/
 *   │   └── SKILL.md
 *   └── db-migrate/
 *       └── SKILL.md
 * </pre>
 *
 * <h2>SKILL.md 文件格式</h2>
 * <pre>
 *   ---
 *   name: code-review
 *   description: Use when reviewing code changes
 *   allowed-tools: [Read, Grep, Glob, Bash(git diff:*)]
 *   context: inline
 *   paths: ["src/**"]
 *   ---
 *
 *   # Code Review Skill
 *   这里是给 LLM 的具体指令...
 * </pre>
 *
 * @author demo — 纯 Java，零依赖，直接运行
 * @see docs/learning/60-skill-learning-guide.md
 */
public class SkillEngineDemo {

    // ============================================================
    // 1. 核心数据结构
    // ============================================================

    /** Skill 定义 —— 从 SKILL.md 解析出来的完整描述 */
    static class SkillDef {
        String name;              // 唯一标识，如 "code-review"
        String description;       // 最重要字段——匹配 + 模型判断都靠它
        String version = "1.0";   // 版本号
        List<String> allowedTools = new ArrayList<>(); // 工具白名单（空=继承默认）
        String context = "inline"; // inline | fork
        List<String> paths = new ArrayList<>(); // glob 条件触发
        String body;              // Markdown 体——发给 LLM 的指令

        @Override
        public String toString() {
            return String.format("Skill[%s] v%s ctx=%s tools=%s paths=%s",
                    name, version, context, allowedTools, paths);
        }
    }

    /** 匹配结果 —— 一个 Skill 是否匹配、匹配得分多少 */
    static class MatchResult {
        final SkillDef skill;
        final int score;          // 匹配得分（越高越相关）
        final String reason;      // 为什么匹配？（用于调试）

        MatchResult(SkillDef skill, int score, String reason) {
            this.skill = skill;
            this.score = score;
            this.reason = reason;
        }

        @Override
        public String toString() {
            return String.format("[%d] %s ← %s", score, skill.name, reason);
        }
    }

    // ============================================================
    // 2. Skill 加载器
    // ============================================================

    /**
     * <h2>加载器核心逻辑</h2>
     * <pre>
     *   1. 递归扫描 skills/ 目录
     *   2. 找到所有 SKILL.md 文件
     *   3. 用 FRONTMATTER_REGEX 提取 YAML 头
     *   4. 简单 YAML 解析（不依赖第三方库）
     *   5. 构造 SkillDef 对象
     *   6. 验证必填字段
     * </pre>
     */
    static class SkillLoader {
        /** 匹配 YAML frontmatter: 开头 --- ... --- */
        private static final Pattern FRONTMATTER = Pattern.compile(
                "^---\\s*\\n([\\s\\S]*?)\\n---\\s*\\n([\\s\\S]*)$");

        /**
         * 递归扫描目录，加载所有 SKILL.md
         */
        List<SkillDef> loadFromDirectory(String skillsPath) throws IOException {
            Path root = Paths.get(skillsPath);
            if (!Files.exists(root)) {
                System.out.println("  [WARN] 目录不存在: " + root.toAbsolutePath());
                return Collections.emptyList();
            }

            List<SkillDef> skills = new ArrayList<>();
            try (var stream = Files.walk(root, 4)) {
                stream.filter(p -> p.getFileName().toString().equals("SKILL.md"))
                        .forEach(p -> {
                            try {
                                SkillDef skill = parseSkillFile(p);
                                if (skill != null) {
                                    skills.add(skill);
                                    System.out.println("  ✅ 加载: " + skill.name
                                            + " (" + p.getParent().getFileName() + ")");
                                }
                            } catch (Exception e) {
                                System.out.println("  ❌ 解析失败: " + p + " — " + e.getMessage());
                            }
                        });
            }

            // 按优先级排序：项目级 .claude/skills/ > 用户级 ~/.claude/skills/ > 内置
            // 简化实现：同名 Skill 取后加载的（后覆盖前）
            Map<String, SkillDef> deduped = new LinkedHashMap<>();
            for (SkillDef s : skills) {
                deduped.put(s.name, s);  // 后者覆盖前者
            }

            return new ArrayList<>(deduped.values());
        }

        /**
         * 解析单个 SKILL.md 文件
         */
        SkillDef parseSkillFile(Path filePath) throws IOException {
            String content = Files.readString(filePath);
            Matcher m = FRONTMATTER.matcher(content);

            if (!m.find()) {
                throw new IOException("No valid frontmatter found (need --- ... ---)");
            }

            String yamlBlock = m.group(1);
            String markdownBody = m.group(2).trim();

            SkillDef skill = new SkillDef();
            skill.body = markdownBody;

            // 简单 YAML 解析（不用 SnakeYAML，保持零依赖）
            parseYamlFrontmatter(yamlBlock, skill);

            // 验证必填字段
            if (skill.name == null || skill.name.isEmpty()) {
                throw new IOException("Missing required field: name");
            }
            if (skill.description == null || skill.description.isEmpty()) {
                throw new IOException("Missing required field: description");
            }

            return skill;
        }

        /**
         * <h2>简易 YAML 解析器（~40 行，零依赖）</h2>
         *
         * <p>只处理 Skill 需要的字段类型：
         * <ul>
         *   <li>标量：name: value</li>
         *   <li>多行标量：description: >- 开头</li>
         *   <li>数组：allowed-tools: [a, b, c]</li>
         * </ul>
         *
         * <p>生产环境应该用 SnakeYAML 或 Jackson YAML，这里是教学目的。</p>
         */
        void parseYamlFrontmatter(String yaml, SkillDef skill) {
            String[] lines = yaml.split("\\n");
            String currentKey = null;
            StringBuilder currentValue = new StringBuilder();
            boolean inMultiline = false;

            for (String line : lines) {
                // 跳过空行和纯注释行
                if (line.trim().isEmpty() || line.trim().startsWith("#")) {
                    continue;
                }

                // 检测是否是新的 key:
                int colonIdx = line.indexOf(':');
                boolean isNewKey = !line.startsWith(" ") && !line.startsWith("\t")
                        && colonIdx > 0 && colonIdx < line.length() - 0;

                if (isNewKey && !inMultiline) {
                    // 保存上一个 key
                    if (currentKey != null) {
                        setField(skill, currentKey, currentValue.toString().trim());
                    }
                    currentKey = line.substring(0, colonIdx).trim();
                    String afterColon = line.substring(colonIdx + 1).trim();

                    // 检查是否是多行模式（>- 或 |）
                    if (afterColon.equals(">-") || afterColon.equals("|") || afterColon.equals(">")) {
                        inMultiline = true;
                        currentValue = new StringBuilder();
                    } else {
                        currentValue = new StringBuilder(afterColon);
                        inMultiline = false;
                    }
                } else {
                    // 续行（缩进或多行模式的后续行）
                    if (currentKey != null) {
                        if (currentValue.length() > 0) {
                            currentValue.append(" ");
                        }
                        currentValue.append(line.trim());
                    }
                }
            }
            // 保存最后一个 key
            if (currentKey != null) {
                setField(skill, currentKey, currentValue.toString().trim());
            }
        }

        /** 把解析出来的 key=value 设到 SkillDef 对象上 */
        void setField(SkillDef skill, String key, String value) {
            switch (key) {
                case "name":
                    skill.name = unquote(value);
                    break;
                case "description":
                    skill.description = unquote(value);
                    break;
                case "version":
                    skill.version = unquote(value);
                    break;
                case "context":
                    skill.context = unquote(value);
                    break;
                case "allowed-tools":
                    skill.allowedTools = parseList(value);
                    break;
                case "paths":
                    skill.paths = parseList(value);
                    break;
                // 忽略其他字段（hooks, model, agent 等——Demo 不需要）
            }
        }

        /** 去掉字符串两端的引号 */
        String unquote(String s) {
            if (s == null) return null;
            s = s.trim();
            if ((s.startsWith("\"") && s.endsWith("\"")) ||
                    (s.startsWith("'") && s.endsWith("'"))) {
                return s.substring(1, s.length() - 1);
            }
            return s;
        }

        /** 解析 YAML 数组：[a, b, c] 或 - a \n - b */
        List<String> parseList(String s) {
            s = s.trim();
            List<String> result = new ArrayList<>();
            if (s.startsWith("[")) {
                // [a, b, c] 格式
                String inner = s.substring(1, s.lastIndexOf(']'));
                for (String item : inner.split(",")) {
                    String trimmed = unquote(item.trim());
                    if (!trimmed.isEmpty()) result.add(trimmed);
                }
            } else if (s.startsWith("-")) {
                // - a \n - b 格式（不支持，留空）
            }
            return result;
        }
    }

    // ============================================================
    // 3. Skill 匹配器
    // ============================================================

    /**
     * <h2>匹配器核心逻辑</h2>
     * <pre>
     *   三种匹配方式（按优先级）：
     *
     *   1. 显式调用：用户输入 "/skill-name" → 精确匹配 name
     *      → +100 分（必定激活）
     *
     *   2. description 匹配：用户输入的关键词出现在 Skill 的 description 中
     *      → 每个关键词命中 +10 分
     *      → 这是最常用的匹配方式
     *
     *   3. paths 匹配：当前涉及的文件路径匹配 Skill 的 paths glob
     *      → +50 分
     *      → 简化实现：子串匹配（不是真正的 glob）
     *
     *   匹配得分 ≥ THRESHOLD 的 Skill 被激活。
     * </pre>
     */
    static class SkillMatcher {
        private static final int EXPLICIT_SCORE = 100;
        private static final int PATH_MATCH_SCORE = 50;
        private static final int DESC_KEYWORD_SCORE = 10;
        private static final int THRESHOLD = 15;

        /**
         * 匹配用户输入，返回激活的 Skill 列表（按得分排序）
         *
         * @param skills     所有已加载的 Skill
         * @param userInput  用户的原始输入
         * @param filePaths  当前涉及的文件路径（可选，用于 paths 匹配）
         */
        List<MatchResult> match(List<SkillDef> skills, String userInput, List<String> filePaths) {
            List<MatchResult> results = new ArrayList<>();

            for (SkillDef skill : skills) {
                int score = 0;
                StringBuilder reason = new StringBuilder();

                // 方式 1: 显式调用 /skill-name 或 /skill-name
                if (userInput.contains("/" + skill.name) || userInput.contains("/ " + skill.name)) {
                    score += EXPLICIT_SCORE;
                    reason.append("显式调用");
                }

                // 方式 2: description 关键词匹配
                // 从 description 和用户输入中提取关键词
                Set<String> descWords = extractKeywords(skill.description);
                Set<String> inputWords = extractKeywords(userInput);

                // 计算重叠的关键词
                Set<String> overlap = new HashSet<>(descWords);
                overlap.retainAll(inputWords);

                if (!overlap.isEmpty()) {
                    score += overlap.size() * DESC_KEYWORD_SCORE;
                    if (reason.length() > 0) reason.append(" + ");
                    reason.append("关键词: ").append(overlap);
                }

                // 方式 3: paths 匹配（如果提供了文件路径）
                if (filePaths != null && !skill.paths.isEmpty()) {
                    for (String fp : filePaths) {
                        for (String glob : skill.paths) {
                            if (simpleGlobMatch(fp, glob)) {
                                score += PATH_MATCH_SCORE;
                                if (reason.length() > 0) reason.append(" + ");
                                reason.append("路径: ").append(fp);
                                break;
                            }
                        }
                    }
                }

                if (score >= THRESHOLD) {
                    results.add(new MatchResult(skill, score, reason.toString()));
                }
            }

            // 按得分降序排列
            results.sort((a, b) -> Integer.compare(b.score, a.score));
            return results;
        }

        /**
         * 从文本中提取有意义的关键词
         * 去掉常见的停用词、标点符号
         */
        Set<String> extractKeywords(String text) {
            // 停用词
            Set<String> stopWords = Set.of(
                    "use", "when", "the", "this", "that", "with", "for", "and",
                    "or", "in", "on", "to", "a", "an", "is", "are", "be", "it",
                    "of", "as", "at", "by", "from", "has", "can", "help", "you",
                    "your", "user", "skill", "should", "will"
            );

            // 分词、转小写、去停用词
            return Arrays.stream(text.toLowerCase().split("[\\s,.!?;:()\\[\\]{}]+"))
                    .map(String::trim)
                    .filter(w -> w.length() > 2)
                    .filter(w -> !stopWords.contains(w))
                    .collect(Collectors.toSet());
        }

        /**
         * 简单的 glob 匹配（教学版，不是完整实现）
         * 只处理 ** 和 * 两种通配符
         */
        boolean simpleGlobMatch(String filePath, String glob) {
            // 去掉 ** 匹配任意路径
            String regex = glob
                    .replace(".", "\\.")
                    .replace("**", "___DOUBLESTAR___")
                    .replace("*", "[^/]*")
                    .replace("___DOUBLESTAR___", ".*");
            return Pattern.compile(regex).matcher(filePath).find();
        }
    }

    // ============================================================
    // 4. Prompt 构建器
    // ============================================================

    /**
     * <h2>注入器核心逻辑</h2>
     * <pre>
     *   设计原则：不破坏 prompt cache！
     *
     *   Claude Code 的做法：
     *   ┌──────────────────────────┐
     *   │ Base System Prompt       │ ← 前缀不变，缓存命中
     *   ├──────────────────────────┤
     *   │ [Active Skills]          │ ← 动态追加在尾部
     *   │ Skill A instructions     │
     *   │ Skill B instructions     │
     *   ├──────────────────────────┤
     *   │ User Messages            │ ← 每次不同
     *   └──────────────────────────┘
     * </pre>
     */
    static class PromptBuilder {

        /**
         * 构建完整的 System Prompt（基础 + 激活的 Skill）
         */
        String buildSystemPrompt(String basePrompt, List<MatchResult> activeSkills) {
            StringBuilder sb = new StringBuilder();
            sb.append(basePrompt);

            if (!activeSkills.isEmpty()) {
                sb.append("\n\n");
                sb.append("## Active Skills\n");
                sb.append("The following skills are activated for this task. ");
                sb.append("Follow their instructions carefully.\n\n");

                for (MatchResult mr : activeSkills) {
                    sb.append("### Skill: ").append(mr.skill.name).append("\n");
                    sb.append(mr.skill.body).append("\n\n");
                }
            }

            return sb.toString();
        }

        /**
         * 构建工具列表描述（受 Skill allowed-tools 限制后）
         */
        String buildToolDescription(List<String> defaultTools, List<MatchResult> activeSkills) {
            // 收集所有激活 Skill 的工具限制
            Set<String> restrictedTools = null;
            for (MatchResult mr : activeSkills) {
                if (!mr.skill.allowedTools.isEmpty()) {
                    if (restrictedTools == null) {
                        restrictedTools = new LinkedHashSet<>();
                    }
                    restrictedTools.addAll(mr.skill.allowedTools);
                }
            }

            if (restrictedTools == null) {
                // 没有工具限制，使用默认工具集
                return "Available tools: " + String.join(", ", defaultTools);
            } else {
                // 有限制——只允许列表中的工具
                return "Available tools (restricted by active skills): "
                        + String.join(", ", restrictedTools);
            }
        }
    }

    // ============================================================
    // 5. 演示入口
    // ============================================================

    public static void main(String[] args) throws IOException {
        System.out.println("=".repeat(65));
        System.out.println("  最小化 Skill 引擎 Demo");
        System.out.println("  演示：加载 → 匹配 → 注入 → 限制");
        System.out.println("=".repeat(65));

        // ── 准备 ──
        SkillLoader loader = new SkillLoader();
        SkillMatcher matcher = new SkillMatcher();
        PromptBuilder builder = new PromptBuilder();

        // ── 第 1 步：加载 Skill ──
        System.out.println("\n📁 第 1 步：加载 Skill\n");

        // 使用内建的示例 Skill（不依赖外部文件）
        List<SkillDef> skills = Arrays.asList(
                createCodeReviewSkill(),
                createDeploySkill(),
                createDbMigrateSkill()
        );
        System.out.println("  共加载 " + skills.size() + " 个 Skill");

        // 也可以从外部目录加载
        // List<SkillDef> skills = loader.loadFromDirectory("skills");

        // ── 第 2 步：模拟用户输入 → 匹配 Skill ──
        System.out.println("\n🔍 第 2 步：匹配 Skill\n");

        String[] testInputs = {
                "帮我审查一下 src/main/java/com/example/UserService.java 的代码",
                "把当前分支部署到生产环境",
                "数据库要加一个 user_roles 表",
                "帮我写一个 Python 脚本来处理 CSV 文件"  // 应该不匹配任何 Skill
        };

        for (String input : testInputs) {
            System.out.println("  用户: \"" + input + "\"");
            List<String> filePaths = extractFilePaths(input);
            List<MatchResult> matches = matcher.match(skills, input, filePaths);

            if (matches.isEmpty()) {
                System.out.println("    → ❌ 没有匹配的 Skill（使用默认行为）\n");
            } else {
                for (MatchResult mr : matches) {
                    System.out.println("    → ✅ " + mr);
                }
                System.out.println();
            }
        }

        // ── 第 3 步：展示注入后的 System Prompt ──
        System.out.println("📝 第 3 步：System Prompt 注入效果\n");

        String userInput = "帮我审查一下 src/main/java/com/example/UserService.java 的代码，看看有没有安全问题";
        List<String> filePaths = extractFilePaths(userInput);
        List<MatchResult> matches = matcher.match(skills, userInput, filePaths);

        String basePrompt = "You are a helpful AI coding assistant. "
                + "You have access to tools: Read, Write, Edit, Bash, Grep, Glob.";

        System.out.println("  匹配到的 Skill:");
        for (MatchResult mr : matches) {
            System.out.println("    ✅ " + mr);
        }

        String fullPrompt = builder.buildSystemPrompt(basePrompt, matches);
        String toolDesc = builder.buildToolDescription(
                List.of("Read", "Write", "Edit", "Bash", "Grep", "Glob"),
                matches
        );

        System.out.println("\n  ── 完整 System Prompt ──");
        System.out.println(truncate(fullPrompt, 1200));
        System.out.println("  ... (truncated)");
        System.out.println("\n  ── 工具列表 ──");
        System.out.println("  " + toolDesc);

        // ── 第 4 步：权限控制演示 ──
        System.out.println("\n🔒 第 4 步：allowed-tools 权限控制\n");

        for (SkillDef skill : skills) {
            System.out.println("  Skill \"" + skill.name + "\":");
            if (skill.allowedTools.isEmpty()) {
                System.out.println("    → 无限制（继承默认工具集）");
            } else {
                System.out.println("    → 限制为: " + skill.allowedTools);
                System.out.println("    → 默认工具集中被移除的: "
                        + diff(List.of("Read", "Write", "Edit", "Bash", "Grep", "Glob"),
                                skill.allowedTools));
            }
        }

        // ── 第 5 步：context 模式对比 ──
        System.out.println("\n🔄 第 5 步：inline vs fork 模式\n");

        System.out.println("  inline 模式 (code-review):");
        System.out.println("    → Skill 注入主对话，能访问完整上下文");
        System.out.println("    → 适合：代码审查、调试、重构（需要看到当前代码状态）");
        System.out.println("    → 风险：Skill 指令占用上下文窗口");

        System.out.println("\n  fork 模式 (deploy):");
        System.out.println("    → Skill 在独立子 Agent 中运行，上下文隔离");
        System.out.println("    → 适合：部署、跑测试、发通知（独立任务）");
        System.out.println("    → 优势：不污染主对话、可并行执行");

        // ── 总结 ──
        System.out.println("\n" + "=".repeat(65));
        System.out.println("  Demo 完成！");
        System.out.println("  核心要点：");
        System.out.println("    1. Skill = YAML 配置（给机器）+ Markdown 指令（给模型）");
        System.out.println("    2. description 是匹配的唯一依据");
        System.out.println("    3. allowed-tools 限制 Skill 可用的工具");
        System.out.println("    4. 注入不破坏 prompt cache（追加在尾部）");
        System.out.println("    5. inline 有上下文，fork 有隔离");
        System.out.println("=".repeat(65));
    }

    // ============================================================
    // 6. 辅助方法
    // ============================================================

    /** 从用户输入中提取文件路径 */
    static List<String> extractFilePaths(String input) {
        List<String> paths = new ArrayList<>();
        // 简单的正则匹配
        Pattern p = Pattern.compile("([a-zA-Z0-9_/.-]*\\.[a-zA-Z]+)");
        Matcher m = p.matcher(input);
        while (m.find()) {
            paths.add(m.group(1));
        }
        return paths;
    }

    /** 计算两个列表的差集 */
    static List<String> diff(List<String> all, List<String> subset) {
        Set<String> s = new HashSet<>(subset);
        return all.stream().filter(t -> !s.contains(t)).toList();
    }

    /** 截断字符串（用于演示输出） */
    static String truncate(String s, int maxLen) {
        if (s.length() <= maxLen) return s;
        return s.substring(0, maxLen);
    }

    // ============================================================
    // 7. 内建示例 Skill（教学用）
    // ============================================================

    /** 代码审查 Skill — inline 模式，有路径触发和工具限制 */
    static SkillDef createCodeReviewSkill() {
        SkillDef s = new SkillDef();
        s.name = "code-review";
        s.description = "Use when reviewing code changes for security vulnerabilities, "
                + "performance issues, and code quality. Covers Java Spring Boot projects.";
        s.version = "1.2";
        s.context = "inline";
        s.allowedTools = List.of("Read", "Grep", "Glob", "Bash(git diff:*)", "Bash(git log:*)");
        s.paths = List.of("src/**", "pom.xml");
        s.body = """
                # Code Review Skill

                You are a senior code reviewer. Follow this checklist:

                ## Security Check
                - No hardcoded credentials or API keys
                - SQL queries use parameterized statements
                - Input validation on all user-facing endpoints
                - No path traversal vulnerabilities

                ## Performance Check
                - No N+1 queries
                - Large collections use pagination
                - No blocking I/O in async contexts

                ## Output Format
                For each finding, use:
                - **File**: `path/to/file.java:line`
                - **Severity**: 🔴 Critical / 🟡 Warning / 🟢 Info
                - **Issue**: What's wrong
                - **Fix**: How to fix it
                """;
        return s;
    }

    /** 部署 Skill — fork 模式，严格工具限制 */
    static SkillDef createDeploySkill() {
        SkillDef s = new SkillDef();
        s.name = "deploy";
        s.description = "Use when deploying applications to production or staging "
                + "environments. Handles build, test, and deployment steps.";
        s.version = "1.0";
        s.context = "fork";  // ← fork 模式：独立子 Agent 执行
        s.allowedTools = List.of(
                "Bash(mvn:*)",
                "Bash(docker:*)",
                "Bash(kubectl:*)",
                "Bash(git push:*)",
                "Read"
        );
        s.body = """
                # Deploy Skill

                You deploy applications safely.

                ## Deployment Checklist

                ### Step 1: Pre-deploy Verification
                1. All tests pass: `mvn test`
                2. No uncommitted changes: `git status`
                3. On the correct branch

                ### Step 2: Build
                ```!
                mvn clean package -DskipTests=false
                ```

                ### Step 3: Deploy
                - Staging first, verify, then production
                - Monitor logs for 2 minutes after deploy
                - Have a rollback plan ready

                ### Step 4: Post-deploy
                - Verify health endpoint
                - Check error rates in monitoring
                - Notify the team
                """;
        return s;
    }

    /** 数据库迁移 Skill — inline 模式，最严格工具限制 */
    static SkillDef createDbMigrateSkill() {
        SkillDef s = new SkillDef();
        s.name = "db-migrate";
        s.description = "Use when creating or modifying database schemas, migrations, "
                + "or running database operations.";
        s.version = "1.0";
        s.context = "inline";
        s.allowedTools = List.of("Read", "Write", "Bash(flyway:*)", "Bash(mvn:*:flyway:*)");
        s.paths = List.of("src/main/resources/db/**");
        s.body = """
                # Database Migration Skill

                You create safe database migrations.

                ## Migration Rules

                1. **Always use Flyway** — never modify the schema directly
                2. **One change per migration** — don't bundle multiple schema changes
                3. **Always include a rollback** — every `V__` migration needs a `U__` undo
                4. **Test rollback first** — run migrate → verify → rollback → verify
                5. **Never drop columns** — deprecate first, drop in a separate release

                ## Migration File Naming
                ```
                V{YYYYMMDD}_{HHmm}__{description}.sql
                ```

                ## Example
                ```sql
                -- V20260603_1430__add_user_roles.sql
                ALTER TABLE users ADD COLUMN role VARCHAR(50) NOT NULL DEFAULT 'USER';
                CREATE INDEX idx_users_role ON users(role);
                ```
                """;
        return s;
    }
}
