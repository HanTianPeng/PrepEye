### Token
  - 大模型处理文本的最小单元；文本会经过Tokenizer

### 模型
  - 本质是根据历史token预测下一个token

### Attention(注意力机制):
  - Q: 查询向量，表示当前token想查询的信息
  - K: 键向量，表示每个token可被匹配的特征
  - V: 值向量，token真正携带的语义内容