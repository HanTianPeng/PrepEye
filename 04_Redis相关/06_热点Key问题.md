#### 热点Key特点
  - 读多写少
  - QPS高
  - 单节点CPU被打满
  - 网络带宽打满
  - event loop延迟
  - 其他请求排队
  - 整个节点RI飙升

#### 解决方案
  - 本地缓存
  - 热点Key拆分
  - 多级缓存: Browser + CDN + Nginx Cache + Local Cache + Redis + DB
  - 读写分离:
  - 限流
  - 异步查询