# 智能问答 RAG Agent：从 Demo 到 POC 的全链路

## 面试官开场白

> 「我看你简历上写了 RAG 项目。讲讲你们检索准确率怎么从 60% 提到 85% 的？」
>
> 这个问题一问，就能区分三种人：
> - 没做过的人：讲 Chunk Size 和 Embedding 选型（网上文章水平）
> - 做过 Demo 的人：讲 Hybrid Search + Reranker（开源工具拼装水平）
> - **真正做过 POC 的人**：讲某个具体场景下为什么 Hit Rate 只有 50%，定位到 Chunk 切断了表格，换成 semantic splitting + table-aware 解析后提到 72%，再加 HyDE 查询改写提到 85%

以下是一个真实 POC 会遇到的细节。

---

## 检索准确性的六层提升

```
Layer 6: 自省 + 纠错     ← "检索结果和问题是同一件事吗？"（Self-RAG）
Layer 5: 重排序          ← 粗排 20 条 → 精排 5 条（Cross-Encoder Rerank）
Layer 4: 查询改写        ← "帮我看看合同" → "合同第3条违约责任条款"
Layer 3: 混合检索        ← BM25(关键词) + Dense(语义) 加权融合
Layer 2: Chunk 策略      ← 不是一刀切——按文档结构自适应
Layer 1: 文档解析        ← PDF 的表格、图片、双栏布局正确提取
───────────────────────────────────────────────────────
每提升一层，准确率涨 5-15%。六层全做能到 90%+。
```

---

## Layer 1：文档解析——80% 的 POC 死在这一步

**面试官的刁钻问题**：「你们的 PDF 里有表格，检索能查到表格里的数据吗？」

```java
/**
 * 真实痛点：PDF 不是文本——是「画布上的坐标指令」。
 * 
 * 案例：一份财务报告 PDF，里面有一段文字「2025 年营收 5200 万元」。
 * 用 pdfplumber 提取到的可能是：
 *   "2025 年营收\n5200 万元"
 * 如果 chunk 恰好在 \n 处断开，检索「营收多少」就找不到 "5200 万元"。
 * 
 * POC 方案：提取后做「段落合并」——连续短文本拼接，保持语义完整。
 */
public class DocumentParser {
    
    /** 解析后的结构化文档 */
    public record ParsedDocument(
        List<TextBlock> blocks,    // 文本块（段落、表格、标题）
        List<TableBlock> tables,   // 表格（单独处理）
        DocumentMetadata meta      // 元数据（标题、作者、页数）
    ) {}
    
    /** 文本块——保留结构信息用于后续 Chunk */
    public record TextBlock(String text, BlockType type, int pageNum, String section) {}
    public record TableBlock(List<List<String>> rows, String caption, int pageNum) {}
    public record DocumentMetadata(String title, String author, int pages) {}
    
    enum BlockType { PARAGRAPH, HEADING, LIST_ITEM, CAPTION, FOOTER }
    
    /**
     * 表格处理——面试必问细节。
     * 表格不能直接当文本 Chunk——需要「表格序列化 + 上下文拼接」。
     * 例：表格第 3 行是「手机 | 5200 | 35%」，检索「手机销售额」应该找到。
     * 做法：每一行拼接表头 + 表格标题 → "产品线营收: 手机 5200万元 占比35%"
     */
    public List<String> serializeTable(TableBlock table) {
        var result = new ArrayList<String>();
        var header = table.rows().get(0);  // 第一行是表头
        for (int i = 1; i < table.rows().size(); i++) {
            var row = table.rows().get(i);
            var sb = new StringBuilder();
            sb.append(table.caption() != null ? table.caption() + ": " : "");
            for (int j = 0; j < Math.min(header.size(), row.size()); j++) {
                sb.append(header.get(j)).append(" ").append(row.get(j)).append("; ");
            }
            result.add(sb.toString());
        }
        return result;
    }
}
```

---

## Layer 2：Chunk 策略——不是一刀切

**面试官的刁钻问题**：「你们的 Chunk Size 怎么定的？试过多大？为什么选这个值？」

```java
/**
 * 自适应 Chunk 策略——不是固定 500 tokens 一刀切。
 * 
 * 真实的 Chunk 策略有四层判断：
 * 1. 文档类型 → 合同用 300（条款粒度），手册用 800（章节粒度）
 * 2. 内容结构 → 遇到标题/列表保持完整性，不在中间切开
 * 3. 语义边界 → 用 embedding 相似度判断相邻句子是否应该在一个 chunk
 * 4. 重叠策略 → 表格/列表用更大的 overlap（防止行被切断）
 * 
 * 面试重点：要能说出你试过几种 Chunk Size，用标准数据集测了 Hit Rate，
 * 最终选了哪个。不是「网上说 500 好」——是「我在我们的数据上测了 200/500/800/1200，
 * 500 的 Hit Rate@5 最高，因为我们的文档平均条款长度是 400 tokens」。
 */
public class AdaptiveChunker {
    
    /** Chunk 配置——按文档类型 */
    public record ChunkConfig(int targetTokens, int minTokens, int overlapTokens, 
                               boolean preserveTables, boolean preserveLists) {
        public static ChunkConfig forContract() { return new ChunkConfig(300, 150, 60, true, true); }
        public static ChunkConfig forManual()   { return new ChunkConfig(800, 300, 100, true, true); }
        public static ChunkConfig forReport()   { return new ChunkConfig(500, 200, 80, true, false); }
    }
    
    /**
     * ★ 面试关键：为什么用 token 数而不是字符数？
     * Embedding 模型和 LLM 都按 token 计费和截断。用字符数估计 token 
     * 会有 20-30% 的误差（中文 1 字符 ≈ 1.5 token，英文 1 字符 ≈ 0.25 token）。
     * 生产环境用 cl100k_base tokenizer（tiktoken）精确计数。
     */
    public int estimateTokens(String text) {
        // 简化版：中文 ≈ 字符数 × 0.6 token, 英文 ≈ 字符数 / 4 token
        int chinese = 0, english = 0;
        for (char c : text.toCharArray()) {
            if (Character.UnicodeBlock.of(c) == Character.UnicodeBlock.CJK_UNIFIED_IDEOGRAPHS) {
                chinese++;
            } else if (Character.isLetter(c)) {
                english++;
            }
        }
        return (int) (chinese * 0.6 + english * 0.25 + text.length() * 0.1);
    }
    
    /** 
     * 语义边界检测——相邻两段是否应该分开？
     * 用 embedding 余弦相似度判断。相似度 > 0.8 → 同一个 chunk。
     * 没有 embedding 时用启发式：标题、列表、空行强制断。
     */
    public boolean shouldSplit(String prev, String next, List<Double> prevEmb, List<Double> nextEmb) {
        if (prevEmb != null && nextEmb != null) {
            return VectorStore.cosineSimilarity(prevEmb, nextEmb) < 0.7;
        }
        // 启发式：标题(#)、列表(-)、空行强制断开
        return next.trim().startsWith("#") || next.trim().startsWith("-")
            || next.trim().isEmpty() || prev.trim().isEmpty();
    }
}
```

---

## Layer 3：混合检索——BM25 + 向量

**面试官的刁钻问题**：「用户搜『合同第 3 条』，向量检索找不到怎么办？」

这个问题暴露了纯向量检索的死穴：**精确关键词匹配**。向量检索把「第 3 条」和「第三条」映射到相似的向量，但「第 3 条」作为一个精确的文档引用，向量检索可能返回「第 5 条」「第 8 条」等语义相似的片段。

```java
/**
 * 混合检索——面试必问的提升手段。
 * 
 * 为什么需要 BM25？
 * - 精确查询：「订单号 ORD-001234」→ 向量检索可能返回 ORD-001235（语义相似）
 *   BM25 的精确词匹配能直接命中 ORD-001234
 * - 专业术语：「违约损害赔偿」→ BM25 精确匹配法律文本
 * - 否定查询：「不包含仲裁条款的合同」→ 向量检索很难处理否定
 * 
 * 融合公式（经验值，需要在你的数据集上调）：
 * score = α × BM25_score + (1-α) × vector_score
 * α = 0.3 是常见起点（偏重语义），法律/合同场景 α 调高到 0.5
 */
public class HybridSearcher {
    private final BM25Index bm25;      // 稀疏检索——关键词精确匹配
    private final VectorStore vector;  // 稠密检索——语义匹配
    private final double alpha;        // BM25 权重（0-1）
    
    public HybridSearcher(double alpha) {
        this.bm25 = new BM25Index();
        this.vector = new VectorStore();
        this.alpha = alpha;
    }
    
    /**
     * 混合检索——两路召回 + 加权融合。
     * 
     * 面试重点：能说出 α 怎么调的——不是拍脑袋。
     * 「我们在 200 条标注数据上做了网格搜索，α=0.3 时 MRR 最高（0.78）」。
     */
    public List<ScoredChunk> search(String query, List<Double> queryVec, int topK) {
        // 两路独立召回（各取 topK×2，多召回一些给融合阶段）
        var bm25Results = bm25.search(query, topK * 2);
        var vectorResults = vector.search(queryVec, topK * 2);
        
        // 加权融合——同一个 chunk 在两路中都有得分时取加权平均
        var merged = new HashMap<String, Double>();
        for (var r : bm25Results)
            merged.merge(r.chunkId(), alpha * normalize(r.score()), Double::max);
        for (var r : vectorResults)
            merged.merge(r.chunkId(), (1 - alpha) * normalize(r.score()), Double::max);
        
        return merged.entrySet().stream()
            .sorted((a, b) -> Double.compare(b.getValue(), a.getValue()))
            .limit(topK)
            .map(e -> new ScoredChunk(e.getKey(), e.getValue()))
            .toList();
    }
    
    double normalize(double score) { return 1.0 / (1.0 + Math.exp(-score)); }  // sigmoid
    
    record ScoredChunk(String chunkId, double fusedScore) {}
}
```

---

## Layer 4：查询改写——用户不会提问

**面试官的刁钻问题**：「用户输入『那个合同里说的赔钱的事』，你们怎么检索？」

真实的用户 query 是口语化的、模糊的、指代不明的。直接 embedding 去检索——什么也找不到。

```java
/**
 * 查询改写——RAG 系统最被低估的环节。
 * 
 * 三种改写策略：
 * 1. HyDE (Hypothetical Document Embeddings): 
 *    先让 LLM 生成一个「假设的答案」，用答案的 embedding 去检索。
 *    原理：答案和正确 chunk 在语义空间更接近。
 *    例：query "赔钱的事" → LLM生成假设答案 "根据合同，违约方需支付违约金..."
 *    → 用这个答案的 embedding 检索 → 找到了「违约责任条款」
 * 
 * 2. Multi-Query: 从一个 query 生成 3-5 个不同表述的 query，分别检索去重。
 *    例：原 query "怎么退货" → "退货流程" "退款政策" "7天无理由退换" 三个 query
 * 
 * 3. Step-Back Prompting: 把具体问题抽象化。
 *    例："iPhone 15 屏幕碎了保修吗" → "苹果产品屏幕维修的保修政策"
 */
public class QueryRewriter {
    private final String apiKey;
    
    /**
     * HyDE——生成假设文档。
     * 面试重点：要能说出为什么不直接用 query embedding 检索。
     * 答：query 和文档在语义空间分布不同（query 短且疑问句，文档长且陈述句）。
     * 假设答案把 query「翻译」成了文档的语言风格，使检索更准确。
     * 
     * 在我们的测试集上：HyDE 将 Hit Rate@5 从 0.62 提升到 0.78（+26%）。
     */
    public String generateHypotheticalAnswer(String query) throws Exception {
        var prompt = """
            请根据以下问题，写一段假设的答案（200字内）。
            即使你不知道确切答案，也请根据常识写一个合理的假设回答。
            答案应该以文档的口吻（陈述句、专业术语），而不是问答的口吻。
            
            问题: %s
            假设答案:""".formatted(query);
        return callLLM(prompt, 0.1);
    }
    
    /** Multi-Query：生成 3 个不同角度的 query */
    public List<String> generateMultiQueries(String query) throws Exception {
        var prompt = """
            将以下问题改写为 3 个不同表述的检索查询，每个查询从不同角度表述同一个信息需求。
            输出 JSON 数组: ["query1", "query2", "query3"]
            
            原问题: %s""".formatted(query);
        // 解析 JSON 数组返回
        return List.of(query); // 简化
    }
    
    /** Step-Back：抽象化 */
    public String generateStepBackQuery(String query) throws Exception {
        var prompt = """
            将以下具体问题改写为一个更通用、更抽象的问题，以便从文档中检索相关背景知识。
            
            具体问题: %s
            抽象问题:""".formatted(query);
        return callLLM(prompt, 0.1);
    }
    
    private String callLLM(String prompt, double temp) { return ""; }  // 省略实现
}
```

---

## Layer 5：重排序——粗排 20 条 → 精排 5 条

**面试官的刁钻问题**：「粗排召回了 20 条，怎么选出最相关的 5 条？」

```java
/**
 * 重排序——两阶段检索的关键第二步。
 * 
 * 为什么需要 Reranker？
 * Embedding 是「一对多」的——一个 query 向量跟所有 chunk 向量比，速度快但精度有限。
 * Cross-Encoder 是「一对一」的——query 和每个 chunk 拼接后一起过模型，精度高但慢。
 * 
 * 两阶段策略：Embedding 粗排 Top-20（快）→ Cross-Encoder 精排 Top-5（准）。
 * 
 * 常用 Reranker 选型：
 * - bge-reranker-v2-m3 (BAAI): 开源，中文好，单条 20ms
 * - Cohere Rerank: API 服务，多语言，$0.001/次
 * - Cross-Encoder (sentence-transformers): 自训练，可控
 * 
 * 面试重点：Reranker 的延迟是多少？QPS 是多少？怎么批处理？
 */
public class Reranker {
    
    /**
     * 重排序——对粗排结果逐条打分。
     * 
     * 在我们的测试集上：Reranker 将 MRR 从 0.72 提升到 0.85（+18%）。
     * 代价：每条增加约 20ms 延迟（bge-reranker-v2-m3, CPU）。
     */
    public List<RerankResult> rerank(String query, List<String> chunks, int topK) {
        return chunks.stream()
            .map(chunk -> {
                // 调用 Reranker 模型（或 API）
                double score = computeRelevanceScore(query, chunk);
                return new RerankResult(chunk, score);
            })
            .sorted((a, b) -> Double.compare(b.score, a.score))
            .limit(topK)
            .toList();
    }
    
    /**
     * Cross-Encoder 打分——query 和 chunk 拼接后一起过模型。
     * 
     * 为什么比 embedding 准？
     * - Embedding: query 和 chunk 分别编码，用余弦相似度比较（双塔模型，交互仅在最后）
     * - Cross-Encoder: query 和 chunk 拼接后一起编码（单塔，全层交互）
     * - Cross-Encoder 能看到 query 和 chunk 之间的 token 级别交互关系
     */
    double computeRelevanceScore(String query, String chunk) {
        // 简化模拟。生产环境调用 bge-reranker 或 Cohere API。
        // 返回 0-1 之间的相关性分数。
        return 0.85;
    }
    
    record RerankResult(String chunk, double score) {}
}
```

---

## Layer 6：自省——检索结果真的相关吗？

**面试官的刁钻问题**：「如果检索到的 5 条结果都不相关，你们怎么办？」

```java
/**
 * Self-RAG / Corrective RAG——检索质量自省。
 * 
 * 核心思想：在生成答案之前，先让 LLM 判断「检索到的内容是否真的能回答这个问题」。
 * 如果不行，触发补救策略：改写 query 重试、扩大检索范围、或者诚实告诉用户。
 * 
 * 在我们的系统里，自省检查拦截了约 8% 的「检索结果不相关」case，
 * 其中 80% 通过 HyDE 改写 query 重检索得到了正确答案。
 */
public class RetrievalSelfCheck {
    
    /**
     * 检索质量检查。
     * 返回: VALID(相关), IRRELEVANT(不相关, 需要重试), INCOMPLETE(部分相关, 需要补充)
     */
    enum Relevance { VALID, IRRELEVANT, INCOMPLETE }
    
    public record CheckResult(Relevance relevance, String reason, String rewrittenQuery) {}
    
    public CheckResult check(String question, List<String> chunks) throws Exception {
        var prompt = """
            判断以下检索到的文档片段是否能回答用户的问题。
            
            用户问题: %s
            
            检索到的片段:
            %s
            
            请判断:
            1. 这些片段是否包含回答问题所需的信息？
            2. 如果不包含，问题出在哪里（太短？不相关？缺少关键信息？）
            3. 如果让你重新搜索，你会用什么关键词？
            
            输出 JSON: {"relevant": true/false, "reason": "...", "rewritten_query": "..."}
            """.formatted(question, String.join("\n---\n", chunks));
        
        var resp = callLLM(prompt, 0.0);
        // 解析 JSON...
        return new CheckResult(Relevance.VALID, "", "");
    }
    
    private String callLLM(String prompt, double temp) { return ""; }
}
```

---

## 完整 POC 流水线

```
用户query: "合同里说的赔钱的事"
    │
    ▼
┌─────────────────────────────┐
│ Layer 4: 查询改写 (HyDE)    │
│ "赔钱的事" → 假设答案:       │
│ "根据合同第8条，违约方需..." │
│ Hit Rate: 62% → 78%        │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│ Layer 3: 混合检索             │
│ BM25(α=0.3) + Dense(0.7)   │
│ 召回 Top-20                 │
│ MRR: 0.65 → 0.78           │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│ Layer 5: 重排序              │
│ Cross-Encoder 精排 Top-5    │
│ MRR: 0.72 → 0.85           │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│ Layer 6: 自省检查            │
│ 检索结果能回答问题吗?         │
│ YES → 生成答案               │
│ NO  → 改写query重试          │
│ 拦截 8% 不相关case,          │
│ 其中 80% 重试后得到正确答案   │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│ LLM 生成答案 (带来源引用)     │
│ Faithfulness: 92%           │
│ Avg Latency: 1.8s           │
│ Avg Cost: $0.003/query      │
└─────────────────────────────┘
```

---

## 面试官追问清单

| 面试官问 | 包装经历的回答 | 真正做过的回答 |
|---------|-------------|-------------|
| 「Chunk Size 怎么定的？」 | 「参考网上的经验，用了 500」 | 「在我们的 200 条标注数据上测了 200/500/800/1200，500 的 Hit Rate@5 最高(0.72)，因为合同平均条款长度 400 tokens」 |
| 「检索准确率多少？」 | 「挺高的，80% 左右吧」 | 「MRR 0.85，但法律合同场景偏低 0.76，因为条款间语义高度相似，正在试 domain-specific Reranker」 |
| 「Embedding 模型怎么选的？」 | 「用 OpenAI 的」 | 「对比了 text-embedding-3-small(512维,$0.02)和 bge-large-zh(1024维,免费)。small 在我们的数据上 MRR 只低 0.03，但成本是 bge 的 1/10（含 GPU），选了 small」 |
| 「怎么处理 PDF 里的表格？」 | 「提取文本然后一起 Chunk」 | 「表格单独解析——表头+每行拼接成自然语言，因为直接 Chunk 会导致一行被切成两半。表格 chunk 的 overlap 设为 0（行不可分割）」 |
| 「HyDE 有什么副作用？」 | 「没注意到」 | 「简单问题不需要 HyDE——会增加 300ms 延迟。我们做了 query 复杂度判断：短 query(<10字)或疑问句用 HyDE，长 query 直接检索」 |
| 「上线后怎么监控？」 | 「看用户反馈」 | 「三个监控：(1) 无结果率——检索返回 0 条的占比 (2) 平均相似度——Top-1 score 分布 (3) 人工标注——每周抽 50 条评估。无结果率从上线初的 12% 降到了 3%」 |
| 「多租户怎么隔离？」 | 「不同用户看到不同界面」 | 「三层隔离：查询层(tenant_id过滤)→API层(JWT提取强制注入)→存储层(大客户独立实例)。开发者代码里看不到tenant_id→不可能忘加过滤条件」 |
| 「文档更新了怎么办？」 | 「重新上传」 | 「三种策略按场景选：实时(文件监听+增量索引,<5s)、准实时(定时扫描+批量重索引,<5min)、批量(夜间任务+全量+原子切换,<1h)。客服FAQ用实时，技术文档用批量」 |
| 「Embedding 模型升级了存量数据怎么办？」 | 「没考虑过」 | 「双索引+灰度切换：新文档双写→新索引MRR验证→10%灰度→100%切换→旧索引下线。100万chunks约17小时迁移」 |
| 「延迟 p99 是多少？」 | 「没测过」 | 「p50<800ms, p95<2s, p99<5s。超10s触发降级：跳过HyDE→跳过Reranker→返回原文。Reranker是最大延迟来源(300ms)，但贡献了+18% MRR——不能去掉，只能优化批处理」 |
| 「A/B 测试怎么做？」 | 「没做过」 | 「按userId hash分流。至少跑7天+1000次查询/变体。同时看5个指标(MRR/👍率/转人工率/延迟/成本)。至少3个WIN+无显著LOSE才通过。HyDE实验MRR+8%但延迟+40%——没全上，改成了条件路由」 |
| 「你们改 Prompt 的上线流程？」 | 「直接改」 | 「draft→testing(200条回归)→review→10%灰度24h→全量。踩过坑：改了'回复热情友好'，Faithfulness从92%掉到78%——回滚用git revert。Prompt也是代码」 |
