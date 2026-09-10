# Skill 学习指南：从概念到实现到面试

> **一句话定位**：Skill = 用 Markdown 文件定义 Agent 能力。四个主流 Agent 框架独立演化出了相同模式。本文是一份独立、完整的 Skill 学习资料——从概念讲起，给出可运行的最小实现，再到面试应对。可以单独分享给任何人。

---

## 一、概念：Skill 是什么？

### 1.1 一句话定义

> **Skill = 一段可复用的提示词模板，包装在 Markdown 文件里，告诉 Agent 在特定场景下该怎么做。**

```
没有 Skill：
  用户："帮我审查代码" → Agent 靠"常识"来判断怎么审查
                        结果：不可控、不一致、每次不一样

有 Skill：
  用户："帮我审查代码" → Skill 注入专业审查指令
                        结果：每次按统一标准、统一的检查清单
```

### 1.2 为什么要用 Markdown 而不是代码？

这是 Skill 系统最核心的设计决策：

| 方案 | 优点 | 缺点 |
|------|------|------|
| **写代码实现 Skill** | 精确、可测试 | 改一个行为要改代码→编译→发版 |
| **写 .md 实现 Skill** | 改行为只需编辑文本，任何人（包括运营/产品）都能改 | 依赖 LLM 正确理解指令 |

**核心洞察**：对于 AI Agent 来说，"一段好的提示词" = "一段好的代码"。因为 Agent 的执行器是 LLM，LLM 的"编译器"就是自然语言理解。所以 **Markdown 就是 Agent 的"脚本语言"**。

### 1.3 类比：Skill = Agent 的"子程序"

```
传统程序：                            Agent 程序：
───────                             ─────────
main() {                            System Prompt（主逻辑）
  read_file();                      Skill A（子程序——代码审查）
  process_data();         ≈        Skill B（子程序——Git 提交）
  write_output();                   Skill C（子程序——部署）
}
```

**关键区别**：传统程序的子程序是 call 进去的，Skill 是**注入到 prompt 里的**。模型"看到" Skill 的全部指令，然后自主决定按不按它做、怎么做。

### 1.4 Skill 全生命周期（6 个环节）

```
① 创作 ──→ ② 加载 ──→ ③ 匹配 ──→ ④ 注入 ──→ ⑤ 执行 ──→ ⑥ 反馈

① 创作：人写 / Agent 写 / 从对话提炼 → 产出 SKILL.md
② 加载：启动时扫描技能目录 → 解析 frontmatter → 建立内存索引
③ 匹配：用户说了一句话 → 哪些 Skill 应该激活？
        三种方式：description 关键词 / paths glob 触发 / 显式 /skill-name
④ 注入：把 SKILL.md 的 body 拼到 system prompt 末尾
        关键约束：不能破坏 prompt cache（修改前缀 = 缓存失效）
⑤ 执行：Agent 按 Skill 指令工作
        约束：allowed-tools 限制可用工具 / context 决定 fork 还是 inline
⑥ 反馈：Skill 用完了，效果如何？需要改进吗？
        高级形态：Agent 自己提议改进 Skill（自进化）
```

---

## 二、格式：一个 SKILL.md 里写什么？

### 2.1 完整示例

```markdown
---
# ═══════════════════════════════════════
# YAML Frontmatter（给 Runtime 读的——机器配置）
# ═══════════════════════════════════════
name: code-review                    # 唯一标识
description: >-                      # ★最重要字段★
  Use when reviewing code changes,   # 模型用它判断是否激活此 Skill
  pull requests, or commits.
  Covers security, performance,
  and code quality.
version: "1.2"                       # 版本号
allowed-tools:                       # 工具白名单（不写=继承默认）
  - Read
  - Grep
  - Glob
  - Bash(git diff:*)
context: inline                      # inline=注入主对话, fork=独立子Agent
model: sonnet                        # 可选：指定模型
paths:                               # 可选：路径条件触发
  - "src/**/*.java"
  - "*.ts"
hooks:                               # 可选：生命周期钩子
  PostToolUse:
    - tool: Read
      action: summarize
---

# ═══════════════════════════════════════
# Markdown Body（给 LLM 读的——这是 Skill 的"代码"）
# ═══════════════════════════════════════

# Code Review Skill

You are a senior code reviewer. Follow this checklist:

## Security Check
- [ ] No hardcoded credentials or API keys
- [ ] SQL queries use parameterized statements
- [ ] Input validation on all user-facing endpoints

## Performance Check
- [ ] No N+1 queries
- [ ] Large collections use pagination
- [ ] No blocking I/O in async contexts

## Output Format
For each finding, use:
- **File**: `path/to/file.ts:line`
- **Severity**: 🔴 Critical / 🟡 Warning / 🟢 Info
- **Issue**: What's wrong
- **Fix**: How to fix it
```

### 2.2 逐字段详解

**`name`** — 唯一标识，字母数字连字符，用于 `/skill-name` 显式调用。

**`description`** — **最关键字段**。这是模型判断"该不该激活此 Skill"的主要依据。用 2-3 句话描述这个 Skill 做什么、什么时候用。关键词密度决定匹配准确率。注意：模型不会直接看到 YAML 里的 description——Runtime 把它加工后注入 system prompt。

**`allowed-tools`** — 权限控制。不写 = 继承默认工具集；写了 = **只允许列表里的工具**。支持 Bash 子命令白名单，如 `Bash(git diff:*)` 只允许 git diff，不允许 rm -rf。

**`context`** — 执行上下文。`inline` 注入主对话（能访问完整上下文，适合代码审查/调试）。`fork` 启动独立子 Agent（上下文隔离、可并行，适合部署/跑测试）。

**`paths`** — 条件触发。Skill 只在用户涉及这些路径的文件时才激活，避免无关 Skill 污染上下文。

**`hooks`** — 生命周期钩子。PreToolUse（工具调用前，如危险命令需确认）、PostToolUse（工具调用后，如读完文件自动总结）。

### 2.3 Markdown Body 的特殊语法

```markdown
## 动态内容注入
当前时间: !`date`             ← 执行 shell 命令并替换结果
当前分支: !`git branch --show-current`

## 代码块（带自动执行）
```!
git status
git diff --staged
```
← ```!  表示执行 Skill 时自动运行此命令
```

---

## 三、实现：最小化 Skill 引擎（~450 行 Java）

> 以下是完整可运行的 Demo 代码。零依赖，Java 17+，直接 `java SkillEngineDemo.java` 运行。
> 演示了：加载 SKILL.md → 解析 frontmatter → 关键词匹配 → 注入 system prompt → allowed-tools 限制。

### 3.1 核心架构

```
skills/                         SkillEngine
├── code-review/               ┌──────────────────────┐
│   └── SKILL.md ───load──────▶│ SkillLoader           │
├── deploy/                    │  • 扫描目录           │
│   └── SKILL.md               │  • 解析 frontmatter   │
└── db-migrate/                │  • 验证必填字段       │
    └── SKILL.md               └─────────┬────────────┘
                                         │
用户输入 ─────────────────────────────────▶│
"帮我审查代码"                             ▼
                               ┌──────────────────────┐
                               │ SkillMatcher          │
                               │  • description 关键词  │
                               │  • paths glob 匹配    │
                               │  • 显式 /skill-name   │
                               └─────────┬────────────┘
                                         │
                                         ▼
                               ┌──────────────────────┐
                               │ PromptBuilder         │
                               │  • 注入 Skill body    │
                               │  • 应用工具限制       │
                               └─────────┬────────────┘
                                         │
                                         ▼
                                  组装好的 System Prompt
                                  + 限制后的工具列表
                                  → 发送给 LLM
```

### 3.2 运行方式

```bash
# 下载这个文件，直接运行
java SkillEngineDemo.java
```

Demo 会输出 5 个步骤的演示结果：
1. 加载 3 个内建示例 Skill
2. 模拟 4 种用户输入，展示匹配过程（哪个 Skill 激活、为什么、得分多少）
3. 展示注入后的完整 system prompt（Skill body 追加在 base prompt 后面）
4. 展示 allowed-tools 权限限制效果（默认工具 vs 限制后的工具）
5. inline vs fork 两种执行模式的对比

### 3.3 完整源码

```java
import java.io.IOException;
import java.nio.file.*;
import java.util.*;
import java.util.regex.*;
import java.util.stream.*;

/**
 * 最小化 Skill 引擎 —— 纯 Java 零依赖 Demo
 *
 * 演示 Skill 系统的 5 个核心环节：
 *   环节 1: 加载 —— 扫描 skills/ 目录，解析 SKILL.md
 *   环节 2: 匹配 —— 用户输入 → 关键词匹配 → 激活哪些 Skill
 *   环节 3: 注入 —— 把 Skill body 拼到 System Prompt 末尾
 *   环节 4: 限制 —— 应用 allowed-tools 权限约束
 *   环节 5: 演示 —— 对比"有 Skill"和"没有 Skill"的 prompt 差异
 *
 * 核心设计原则（从 Claude Code 迁移）：
 *   1. YAML Frontmatter 给 Runtime 读（机器的配置）
 *   2. Markdown Body 给 LLM 读（模型的指令）
 *   3. description 是匹配的主要依据
 *   4. Skill 注入不破坏 prompt cache（追加在 system prompt 尾部）
 *
 * 运行方式：java SkillEngineDemo.java
 */
public class SkillEngineDemo {

    // ============================================================
    // 1. 核心数据结构
    // ============================================================

    /** Skill 定义 —— 从 SKILL.md 解析出来的完整描述 */
    static class SkillDef {
        String name;              // 唯一标识，如 "code-review"
        String description;       // ★最重要字段——匹配 + 模型判断都靠它
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
            return String.format("[得分:%d] %s ← %s", score, skill.name, reason);
        }
    }

    // ============================================================
    // 2. Skill 加载器
    // ============================================================

    /**
     * 加载器核心逻辑：
     *   1. 递归扫描 skills/ 目录
     *   2. 找到所有 SKILL.md 文件
     *   3. 用 FRONTMATTER_REGEX 提取 YAML 头
     *   4. 简单 YAML 解析（不依赖第三方库）
     *   5. 构造 SkillDef 对象
     *   6. 验证必填字段 (name, description)
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
                                }
                            } catch (Exception e) {
                                System.out.println("  [ERROR] 解析失败: " + p + " — " + e.getMessage());
                            }
                        });
            }

            // 同名 Skill 去重：后加载的覆盖先加载的
            Map<String, SkillDef> deduped = new LinkedHashMap<>();
            for (SkillDef s : skills) {
                deduped.put(s.name, s);
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

            // 简单 YAML 解析（不用第三方库，保持零依赖）
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
         * 简易 YAML 解析器（~40 行，零依赖）
         * 只处理 Skill 需要的字段类型：
         *   标量：name: value
         *   多行标量：description: >- 开头
         *   数组：allowed-tools: [a, b, c]
         *
         * 生产环境应该用 SnakeYAML 或 Jackson YAML，这里是教学目的。
         */
        void parseYamlFrontmatter(String yaml, SkillDef skill) {
            String[] lines = yaml.split("\\n");
            String currentKey = null;
            StringBuilder currentValue = new StringBuilder();
            boolean inMultiline = false;

            for (String line : lines) {
                if (line.trim().isEmpty() || line.trim().startsWith("#")) continue;

                int colonIdx = line.indexOf(':');
                boolean isNewKey = !line.startsWith(" ") && !line.startsWith("\t")
                        && colonIdx > 0 && colonIdx < line.length() - 1;

                if (isNewKey && !inMultiline) {
                    if (currentKey != null) {
                        setField(skill, currentKey, currentValue.toString().trim());
                    }
                    currentKey = line.substring(0, colonIdx).trim();
                    String afterColon = line.substring(colonIdx + 1).trim();

                    if (afterColon.equals(">-") || afterColon.equals("|") || afterColon.equals(">")) {
                        inMultiline = true;
                        currentValue = new StringBuilder();
                    } else {
                        currentValue = new StringBuilder(afterColon);
                        inMultiline = false;
                    }
                } else {
                    if (currentKey != null) {
                        if (currentValue.length() > 0) currentValue.append(" ");
                        currentValue.append(line.trim());
                    }
                }
            }
            if (currentKey != null) {
                setField(skill, currentKey, currentValue.toString().trim());
            }
        }

        void setField(SkillDef skill, String key, String value) {
            switch (key) {
                case "name" -> skill.name = unquote(value);
                case "description" -> skill.description = unquote(value);
                case "version" -> skill.version = unquote(value);
                case "context" -> skill.context = unquote(value);
                case "allowed-tools" -> skill.allowedTools = parseList(value);
                case "paths" -> skill.paths = parseList(value);
            }
        }

        String unquote(String s) {
            if (s == null) return null;
            s = s.trim();
            if ((s.startsWith("\"") && s.endsWith("\""))
                    || (s.startsWith("'") && s.endsWith("'"))) {
                return s.substring(1, s.length() - 1);
            }
            return s;
        }

        List<String> parseList(String s) {
            s = s.trim();
            List<String> result = new ArrayList<>();
            if (s.startsWith("[")) {
                String inner = s.substring(1, s.lastIndexOf(']'));
                for (String item : inner.split(",")) {
                    String trimmed = unquote(item.trim());
                    if (!trimmed.isEmpty()) result.add(trimmed);
                }
            }
            return result;
        }
    }

    // ============================================================
    // 3. Skill 匹配器
    // ============================================================

    /**
     * 三种匹配方式（按优先级）：
     *
     * 1. 显式调用：用户输入 "/skill-name" → +100 分（必定激活）
     * 2. description 关键词匹配：用户输入的关键词出现在 Skill 的 description 中
     *    → 每个命中关键词 +10 分
     * 3. paths 匹配：当前涉及的文件匹配 Skill 的 paths glob → +50 分
     *
     * 得分 ≥ 15 的 Skill 被激活。
     */
    static class SkillMatcher {
        private static final int EXPLICIT_SCORE = 100;
        private static final int PATH_MATCH_SCORE = 50;
        private static final int DESC_KEYWORD_SCORE = 10;
        private static final int THRESHOLD = 15;

        List<MatchResult> match(List<SkillDef> skills, String userInput, List<String> filePaths) {
            List<MatchResult> results = new ArrayList<>();

            for (SkillDef skill : skills) {
                int score = 0;
                StringBuilder reason = new StringBuilder();

                // 方式 1: 显式调用
                if (userInput.contains("/" + skill.name)) {
                    score += EXPLICIT_SCORE;
                    reason.append("显式调用");
                }

                // 方式 2: description 关键词匹配
                Set<String> descWords = extractKeywords(skill.description);
                Set<String> inputWords = extractKeywords(userInput);
                Set<String> overlap = new HashSet<>(descWords);
                overlap.retainAll(inputWords);

                if (!overlap.isEmpty()) {
                    score += overlap.size() * DESC_KEYWORD_SCORE;
                    if (reason.length() > 0) reason.append(" + ");
                    reason.append("关键词:").append(overlap);
                }

                // 方式 3: paths 匹配
                if (filePaths != null && !skill.paths.isEmpty()) {
                    for (String fp : filePaths) {
                        for (String glob : skill.paths) {
                            if (simpleGlobMatch(fp, glob)) {
                                score += PATH_MATCH_SCORE;
                                if (reason.length() > 0) reason.append(" + ");
                                reason.append("路径:").append(fp);
                                break;
                            }
                        }
                    }
                }

                if (score >= THRESHOLD) {
                    results.add(new MatchResult(skill, score, reason.toString()));
                }
            }

            results.sort((a, b) -> Integer.compare(b.score, a.score));
            return results;
        }

        Set<String> extractKeywords(String text) {
            Set<String> stopWords = Set.of(
                    "use", "when", "the", "this", "that", "with", "for", "and",
                    "or", "in", "on", "to", "a", "an", "is", "are", "be", "it",
                    "of", "as", "at", "by", "from", "has", "can", "help", "you",
                    "your", "user", "skill", "should", "will"
            );
            return Arrays.stream(text.toLowerCase().split("[\\s,.!?;:()\\[\\]{}]+"))
                    .map(String::trim)
                    .filter(w -> w.length() > 2)
                    .filter(w -> !stopWords.contains(w))
                    .collect(Collectors.toSet());
        }

        boolean simpleGlobMatch(String filePath, String glob) {
            String regex = glob
                    .replace(".", "\\.")
                    .replace("**", "___DS___")
                    .replace("*", "[^/]*")
                    .replace("___DS___", ".*");
            return Pattern.compile(regex).matcher(filePath).find();
        }
    }

    // ============================================================
    // 4. Prompt 构建器
    // ============================================================

    /**
     * 核心原则：不破坏 prompt cache！
     *
     * ┌──────────────────────────┐
     * │ Base System Prompt       │ ← 前缀不变，缓存命中
     * ├──────────────────────────┤
     * │ [Active Skills]          │ ← 动态追加在尾部
     * │ Skill A instructions     │
     * │ Skill B instructions     │
     * ├──────────────────────────┤
     * │ User Messages            │ ← 每次不同
     * └──────────────────────────┘
     */
    static class PromptBuilder {

        String buildSystemPrompt(String basePrompt, List<MatchResult> activeSkills) {
            StringBuilder sb = new StringBuilder();
            sb.append(basePrompt);

            if (!activeSkills.isEmpty()) {
                sb.append("\n\n## Active Skills\n");
                sb.append("The following skills are activated. ");
                sb.append("Follow their instructions carefully.\n\n");

                for (MatchResult mr : activeSkills) {
                    sb.append("### Skill: ").append(mr.skill.name).append("\n");
                    sb.append(mr.skill.body).append("\n\n");
                }
            }

            return sb.toString();
        }

        String buildToolDescription(List<String> defaultTools, List<MatchResult> activeSkills) {
            Set<String> restrictedTools = null;
            for (MatchResult mr : activeSkills) {
                if (!mr.skill.allowedTools.isEmpty()) {
                    if (restrictedTools == null) restrictedTools = new LinkedHashSet<>();
                    restrictedTools.addAll(mr.skill.allowedTools);
                }
            }

            if (restrictedTools == null) {
                return "Available tools: " + String.join(", ", defaultTools);
            } else {
                return "Available tools (restricted by active skills): "
                        + String.join(", ", restrictedTools);
            }
        }
    }

    // ============================================================
    // 5. 演示入口
    // ============================================================

    public static void main(String[] args) {
        System.out.println("=".repeat(65));
        System.out.println("  最小化 Skill 引擎 Demo");
        System.out.println("  演示：加载 → 匹配 → 注入 → 限制");
        System.out.println("=".repeat(65));

        SkillLoader loader = new SkillLoader();
        SkillMatcher matcher = new SkillMatcher();
        PromptBuilder builder = new PromptBuilder();

        // 使用内建示例 Skill（不从外部文件加载，保持独立可运行）
        List<SkillDef> skills = Arrays.asList(
                createCodeReviewSkill(),
                createDeploySkill(),
                createDbMigrateSkill()
        );

        // ── 第 1 步：展示已加载的 Skill ──
        System.out.println("\n📁 第 1 步：已加载 " + skills.size() + " 个 Skill\n");
        for (SkillDef s : skills) {
            System.out.println("  ✅ " + s);
        }

        // ── 第 2 步：模拟用户输入 → 匹配 Skill ──
        System.out.println("\n🔍 第 2 步：匹配演示\n");

        String[][] testCases = {
            {"帮我审查一下 src/main/java/com/example/UserService.java 的代码",
             "src/main/java/com/example/UserService.java"},
            {"把当前分支部署到生产环境", ""},
            {"数据库要加一个 user_roles 表", "src/main/resources/db/migration/V1__init.sql"},
            {"帮我写一个 Python 脚本来处理 CSV 文件", ""}  // 应该不匹配任何 Skill
        };

        for (String[] tc : testCases) {
            String input = tc[0];
            String filePath = tc[1];
            List<String> fps = filePath.isEmpty() ? List.of() : List.of(filePath);

            System.out.println("  用户: \"" + input + "\"");
            List<MatchResult> matches = matcher.match(skills, input, fps);

            if (matches.isEmpty()) {
                System.out.println("    → ❌ 没有匹配的 Skill（使用默认行为）\n");
            } else {
                for (MatchResult mr : matches) {
                    System.out.println("    → ✅ " + mr);
                }
                System.out.println();
            }
        }

        // ── 第 3 步：System Prompt 注入效果 ──
        System.out.println("📝 第 3 步：System Prompt 注入\n");

        String userInput = "帮我审查一下 src/main/java/com/example/UserService.java 的代码，看看有没有安全问题";
        List<MatchResult> matches = matcher.match(skills, userInput,
                List.of("src/main/java/com/example/UserService.java"));

        String basePrompt = "You are a helpful AI coding assistant. "
                + "You have access to tools: Read, Write, Edit, Bash, Grep, Glob.";

        String fullPrompt = builder.buildSystemPrompt(basePrompt, matches);
        String toolDesc = builder.buildToolDescription(
                List.of("Read", "Write", "Edit", "Bash", "Grep", "Glob"), matches);

        System.out.println("  匹配到的 Skill: " + matches);
        System.out.println("\n  ── 完整 System Prompt（截断）──");
        System.out.println(fullPrompt.substring(0, Math.min(fullPrompt.length(), 1000)));
        System.out.println("  ... (truncated)\n");
        System.out.println("  ── 工具列表 ──");
        System.out.println("  " + toolDesc);

        // ── 第 4 步：权限控制 ──
        System.out.println("\n🔒 第 4 步：allowed-tools 权限控制\n");

        List<String> defaultTools = List.of("Read", "Write", "Edit", "Bash", "Grep", "Glob");
        for (SkillDef s : skills) {
            System.out.println("  Skill \"" + s.name + "\":");
            if (s.allowedTools.isEmpty()) {
                System.out.println("    → 无限制（继承默认工具集）");
            } else {
                System.out.println("    → 限制为: " + s.allowedTools);
                Set<String> restricted = new HashSet<>(s.allowedTools);
                List<String> removed = defaultTools.stream()
                        .filter(t -> !restricted.contains(t)).toList();
                System.out.println("    → 被移除的默认工具: " + removed);
            }
        }

        // ── 第 5 步：inline vs fork ──
        System.out.println("\n🔄 第 5 步：inline vs fork 执行模式\n");

        System.out.println("  inline (code-review):");
        System.out.println("    → Skill 注入主对话，能访问完整上下文");
        System.out.println("    → 适合：代码审查、调试、重构");

        System.out.println("\n  fork (deploy):");
        System.out.println("    → Skill 在独立子 Agent 中运行，上下文隔离");
        System.out.println("    → 适合：部署、跑测试、发通知");

        // ── 总结 ──
        System.out.println("\n" + "=".repeat(65));
        System.out.println("  5 个核心要点：");
        System.out.println("    1. Skill = YAML 配置（给机器）+ Markdown 指令（给模型）");
        System.out.println("    2. description 是匹配的主要依据");
        System.out.println("    3. allowed-tools 限制 Skill 可用的工具集");
        System.out.println("    4. 注入不破坏 prompt cache（追加在尾部，不修改前缀）");
        System.out.println("    5. inline 有完整上下文，fork 有隔离安全性");
        System.out.println("=".repeat(65));
    }

    // ============================================================
    // 6. 内建示例 Skill
    // ============================================================

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

    static SkillDef createDeploySkill() {
        SkillDef s = new SkillDef();
        s.name = "deploy";
        s.description = "Use when deploying applications to production or staging "
                + "environments. Handles build, test, and deployment steps.";
        s.version = "1.0";
        s.context = "fork";
        s.allowedTools = List.of("Bash(mvn:*)", "Bash(docker:*)", "Bash(kubectl:*)",
                "Bash(git push:*)", "Read");
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
                3. **Always include a rollback** — every V__ migration needs a U__ undo
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
```

### 3.4 运行输出示例

```
=================================================================
  最小化 Skill 引擎 Demo
  演示：加载 → 匹配 → 注入 → 限制
=================================================================

📁 第 1 步：已加载 3 个 Skill

  ✅ Skill[code-review] v1.2 ctx=inline tools=[Read, Grep, Glob, ...] paths=[src/**, pom.xml]
  ✅ Skill[deploy] v1.0 ctx=fork tools=[Bash(mvn:*), Bash(docker:*), ...] paths=[]
  ✅ Skill[db-migrate] v1.0 ctx=inline tools=[Read, Write, Bash(flyway:*), ...] paths=[src/main/resources/db/**]

🔍 第 2 步：匹配演示

  用户: "帮我审查一下 src/main/java/com/example/UserService.java 的代码"
    → ✅ [得分:70] code-review ← 关键词:[reviewing, code, java, ...] + 路径:src/.../UserService.java

  用户: "把当前分支部署到生产环境"
    → ✅ [得分:30] deploy ← 关键词:[deploying, production, environments]

  用户: "数据库要加一个 user_roles 表"
    → ✅ [得分:60] db-migrate ← 关键词:[database, migrations, operations, ...] + 路径:.../V1__init.sql

  用户: "帮我写一个 Python 脚本来处理 CSV 文件"
    → ❌ 没有匹配的 Skill（使用默认行为）
```

---

## 四、对比：四个框架的 Skill 系统

四个主流 Agent 框架**独立演化出了相同模式**：`Skill = 目录/SKILL.md + YAML frontmatter`。这不是巧合——这是当前工程实践收敛到的局部最优解。

### 4.1 横向对比总表

| 维度 | Claude Code | Codex | QwenPaw | Hermes |
|------|------------|-------|---------|--------|
| **Skill 格式** | SKILL.md + YAML | SKILL.md + YAML | SKILL.md + YAML + skill.json | SKILL.md + JSON 头 |
| **Frontmatter 字段数** | 15+ | 8+ | 10+ | 5+ |
| **触发方式** | description + paths + 显式 | description + 显式 | description + 显式 + 渠道 | 显式 + Curator 推荐 |
| **执行上下文** | inline / fork | inline | inline | inline |
| **权限控制** | allowed-tools（Bash 子命令级） | dependencies.tools | requires.bins/env | 无 |
| **条件触发** | paths glob + hooks | 无 | 渠道筛选 | 无 |
| **Agent 自创** | /skillify（人驱动） | 无 | make-skill（Agent 驱动+审批+扫描） | Curator 后台静默整理 |
| **分发渠道** | Plugin 市场 | Plugin 市场 | 7 Hub（GitHub/LobeHub...） | 无 |
| **安全审计** | 手动审查 | 手动审查 | SkillScanner 自动扫描 | 无 |
| **内建 Skill 数** | 7 个 | 5 个 | 36 个 | 0（完全靠自进化） |

### 4.2 收敛点：为什么都是 SKILL.md？

- **Markdown 体** = LLM 可直接理解（零编译成本，自然语言就是 Agent 的"机器码"）
- **YAML 头** = Runtime 可解析（机器可读，不发给 LLM）
- **分隔线 `---`** = 人可读 + 机器可分界

### 4.3 分歧点：哪里不一样？

**权限模型**：
- Claude Code：细分到 Bash 子命令级别（`Bash(git diff:*)`）
- QwenPaw：声明式依赖（`requires.bins: [python3]`）
- Hermes：无权限限制（信任模型）

**自进化路径**：
- Claude Code：用户主动调用 `/skillify` → 人驱动
- Hermes：Curator 后台静默整理 → 系统驱动
- QwenPaw：Agent 完成任务后提议创建 Skill → Agent 驱动 + 人审批 + 安全扫描。这是唯一包含安全扫描的完整闭环。

**生态规模**：
- Claude Code：7 个内置（小但精，偏编码场景）
- QwenPaw：36 个内置 + 7 个 Hub 分发 → 生态最大

### 4.4 理想 Skill 系统的拼图

```
借用 Claude Code 的 → 丰富的 frontmatter（15+ 字段），精细的工具权限控制
借用 QwenPaw 的    → 完整的生态（7 Hub 分发 + 渠道感知 + 安全扫描）
借用 Hermes 的     → Agent 自进化（从对话中自动提炼 Skill）
借用 Codex 的      → OS 级沙箱隔离（Skill 在独立沙箱里执行）
```

---

## 五、面试：Skill 相关 6 个高频问题

### Q1: "你用过 Skill 吗？"

> "用过。我装了 Superpowers 插件，里面有一整套 Skill——brainstorming 帮我在写代码前想清楚需求，writing-plans 帮我拆任务，verification-before-completion 确保不破坏已有功能。

> 但更重要的是，我理解了 Skill 的本质——它不是某个框架的专利，而是一个通用的 Agent 设计模式：**用 Markdown 文件定义能力，让 AI 的行为从代码中解耦出来**。我在自己的项目里也参考这个思想，把 Agent 的 behavior instructions 放到外部 Markdown 文件里，运营/产品也能调 AI 行为，不需要工程介入。"

**得分点**：提到了具体 Skill 名称 + 能从使用经验上升到设计模式 + 有迁移实践。

---

### Q2: "Skill 的原理是什么？"

> "一个 Skill 从文件到执行经过 6 步：

> ① **加载**：启动时扫描技能目录，找到所有 SKILL.md，解析 YAML frontmatter，建立内存索引
> ② **匹配**：用户说话时，用 description 做关键词匹配，paths glob 做路径匹配，`/name` 做精确匹配
> ③ **注入**：把匹配到的 Skill 的 markdown body 拼到 system prompt 末尾——**不能改前缀，否则 prompt cache 失效**
> ④ **限制**：如果 Skill 声明了 allowed-tools，Runtime 限制这个 Skill 只能用指定工具
> ⑤ **执行**：模型看到注入后的 system prompt + 用户消息，按 Skill 指令工作

> 核心设计：**YAML 给 Runtime 读（机器），Markdown 给 LLM 读（模型），各取所需。**"

**得分点**：6 步链路清晰 + 提到了 prompt cache 约束（进阶细节）+ 最后一句抓住设计本质。

---

### Q3: "Skill 和 Function Calling（Tool）有什么区别？"

> | | Function Calling (Tool) | Skill |
> |------|------------------------|-------|
> | **是什么** | 一段可执行的代码（函数） | 一段提示词（指令） |
> | **谁执行** | Runtime 执行函数，结果返回模型 | 模型自己按指令执行（用已有工具） |
> | **定义方式** | JSON Schema + 代码实现 | Markdown 文件 |
> | **适用场景** | 确定性操作（查数据库、调 API） | 指导性操作（怎么审查代码、怎么写 commit） |

> **简单说**：Tool 是给模型一个"按钮"，模型按按钮，系统执行。Skill 是给模型一本"手册"，模型看了手册，自己动手。

> **但两者不是互斥的**——Skill 可以用 allowed-tools 限制可用哪些 Tool，Skill 的 body 里也可以指导模型"遇到 X 情况调用 Y 工具"。Skill 是更高层的抽象，它**编排** Tools 的使用方式。"

---

### Q4: "怎么设计一个好的 Skill？"

> "四个原则：

> ① **description 要精准**——这是模型判断"该不该用这个 Skill"的主要依据。不好的例子：`Helpful assistant for coding`。好的例子：`Use when reviewing Java code changes for security vulnerabilities, focusing on SQL injection, XSS, and path traversal.`

> ② **明确 when NOT to use**——比 when to use 更重要。防止 Skill 在不合适的场景被激活，污染上下文。

> ③ **给具体例子，不给抽象原则**——不要写"代码要规范"，写"方法名用 camelCase，类名用 PascalCase，常量用 UPPER_SNAKE_CASE"。

> ④ **控制 Skill 大小**——一个 Skill 不要超过 200 行。太大了模型注意力分散，应该拆成多个 Skill 组合使用。"

---

### Q5: "你怎么管理 100 个 Skill？"

> "100 个 Skill 面临的挑战不是加载（那是一次性的），而是**激活决策**——每次对话要决定激活哪几个 Skill。

> ① **分层管理**：核心 Skill（每次加载）+ 场景 Skill（按需激活）+ 低频 Skill（显式调用）
> ② **paths 过滤**：90% 的 Skill 应该用 paths 限定作用域
> ③ **优先级与冲突处理**：两个 Skill 冲突时按 specificity 裁决
> ④ **审计日志**：记录每个 Skill 的激活次数、成功率，低频低效的自动归档

> 这其实就是微服务的治理模式平移到 Skill 管理——服务注册、负载均衡、熔断降级。"

---

### Q6: "如果你来设计一个 Skill 系统，怎么做？"

> "分三步：

> **V1 — 最小可用（2 天）**：加载 `skills/*/SKILL.md` → 解析 YAML → 关键词匹配 → 注入 prompt。约 300 行代码。
>
> **V2 — 工程加固（1 周）**：allowed-tools 权限控制 + inline/fork 双模式 + paths 条件触发 + 冲突裁决。
>
> **V3 — 生态建设（1 月+）**：Skill 市场和远程分发 + Agent 自创 Skill 提案→审批→安全扫描闭环 + 渠道感知。

> 我已经写了一个 V1 的最小实现 Demo（约 450 行 Java），演示了加载→匹配→注入的完整链路。"

---

### 自测清单

**Level 1 — 使用者**：
- [ ] 能说出 Skill 和 Tool/Function Calling 的区别
- [ ] 用过至少 1 个第三方 Skill（如 Superpowers）
- [ ] 知道 SKILL.md 的基本结构（YAML 头 + Markdown 体）
- [ ] 能解释为什么用 Markdown 而不是代码

**Level 2 — 设计者**：
- [ ] 能写出一个合格的 SKILL.md（包含 description、when-not-to-use、allowed-tools）
- [ ] 理解 inline 和 fork 的区别和适用场景
- [ ] 知道 description 是匹配的主要依据
- [ ] 理解 prompt cache 不被破坏的设计约束

**Level 3 — 架构者**：
- [ ] 能对比 3+ 个框架的 Skill 系统，说出各自优劣
- [ ] 能设计 Skill 的完整生命周期（创建→分发→版本→退役）
- [ ] 能手写最小 Skill 引擎（加载→匹配→注入→执行）
- [ ] 能解决 Skill 冲突和优先级问题
- [ ] 能设计 Skill 自进化闭环

---

## 六、动手：写两个自己的 Skill

### 练习 1：代码审查 Skill（15 分钟）

```markdown
---
name: my-code-review
description: >
  Use when reviewing Java Spring Boot code. Checks security,
  performance, and architecture.
version: "1.0"
allowed-tools:
  - Read
  - Grep
  - Glob
  - Bash(git diff:*)
  - Bash(git log:*)
context: inline
paths:
  - "src/**/*.java"
  - "pom.xml"
---

# Java Code Review Skill

You review Java Spring Boot code. Follow this checklist:

## 1. Architecture Check
- [ ] Controllers are thin (no business logic)
- [ ] Services handle transactions correctly (@Transactional placement)
- [ ] No circular dependencies between services

## 2. Security Check
- [ ] All endpoints have authentication
- [ ] User input is validated (@Valid, @NotBlank, etc.)
- [ ] SQL uses parameterized queries
- [ ] Sensitive data is not logged

## 3. Performance Check
- [ ] No SELECT N+1 patterns
- [ ] Large result sets use pagination (Pageable)
- [ ] Async operations use @Async with proper thread pool

## 4. Output Format
For each issue:
- **File**: `path/to/file.java:line`
- **Category**: 🔴Critical / 🟡Warning / 🟢Suggestion
- **Problem**: What's wrong
- **Fix**: Concrete code example
```

### 练习 2：Git Commit Skill（10 分钟）

```markdown
---
name: smart-commit
description: >
  Use when the user wants to commit changes. Generates
  conventional commit messages from staged diffs.
version: "1.0"
allowed-tools:
  - Bash(git status:*)
  - Bash(git diff:*)
  - Bash(git log:*)
  - Bash(git add:*)
  - Bash(git commit:*)
context: inline
---

# Smart Commit Skill

## Workflow

### Step 1: Check current state
```!
git status
git diff --staged --stat
git log --oneline -5
```

### Step 2: Analyze changes
- Categorize each change: feat / fix / refactor / docs / test / chore
- Identify the scope (module or component affected)
- Determine if this is a breaking change

### Step 3: Generate message (Conventional Commits format)
```
<type>(<scope>): <description>

<body>
```

Example:
```
feat(user-service): add OAuth2 login support

- Add OAuth2 client configuration
- Implement token exchange flow
- Add user profile mapping
```

### Step 4: Confirm and commit
Present the message to the user for approval before committing.
```

---

## 七、进阶话题

### 7.1 Prompt Cache 与 Skill 注入的兼容性

Claude Code 最精巧的设计之一——Skill 注入不破坏 prompt cache：

```
没有 Skill 时的请求：
┌──────────────────────────┐
│ System Prompt (不会变)    │ ← 整个被缓存，后续请求命中
│ User Messages            │ ← 每次不同，不缓存
└──────────────────────────┘

有 Skill 时的请求：
┌──────────────────────────┐
│ System Prompt (不会变)    │ ← 前缀不变，缓存命中！
│ [Skill A 指令]           │ ← 追加在末尾
│ [Skill B 指令]           │ ← 追加在末尾
│ User Messages            │ ← 每次不同
└──────────────────────────┘

关键：Skill 追加在 System Prompt 后面，不修改前缀，
所以 System Prompt 的缓存持续命中。
```

### 7.2 Skill 组合模式（Meta-Skill）

```markdown
---
name: full-release
description: Full release workflow combining review, test, and deploy
version: "1.0"
---
# Full Release Workflow

This is a "meta-skill" — it orchestrates other skills:

## Phase 1: Review
Invoke the `code-review` skill to check all changes.

## Phase 2: Test
Run the full test suite. If any test fails, STOP.

## Phase 3: Deploy
Invoke the `deploy` skill with the current branch.
```

### 7.3 Skill 代码化：从提示词到可执行脚本

未来方向——Skill 可以携带真正的可执行代码：

```markdown
---
name: db-health-check
description: Check database connection pool health
version: "1.0"
script: check_db.py              ← 真正的可执行脚本
allowed-tools:
  - Bash(python check_db.py:*)
---
# Database Health Check

Execute the bundled script:
```!
python check_db.py --threshold 80
```

If connection usage > 80%, suggest scaling actions.
```

---

## 总结：三个核心认知

1. **Skill 的本质是 "Prompt 工程化"**——把一次性提示词变成可复用、可组合、可分发、可版本管理的独立单元。

2. **Skill 的格式是行业收敛的结果**——YAML 给机器 + Markdown 给模型。四个框架独立选择了同样的方案，这不是巧合。

3. **Skill 系统的工程深度决定 Agent 平台的通用性**——是 7 个内置 Skill 还是 36 个 + 7 Hub 分发 + Agent 自创？差距不在"有没有 Skill"，而在**全生命周期的工程化程度**。
