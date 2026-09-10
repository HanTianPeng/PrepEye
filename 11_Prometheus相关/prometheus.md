### 架构图:
  - Exporters: 采集层
  - Prometheus Server: 调度层, pull模型(主动拉取)、调度Scrape、解析metrics、生成time series
  - TSDB: 存储层, 本地时序数据库, WAL - Memory - Block(磁盘)
  - PromQL: 查询层, Grafana、Prometheus UI
  - Alerting: 告警层, rule engine、alertmanager

### rate与irate区别:
  - rate是一段时间内的平均增长速率,irate是基于最近两个点的瞬时速率
  - `rate = Δvalue / Δtime`

### Prometheus的数据样例: 多维时间序列数据库
  - `metric{label=value} timestamp value`

### Label是什么:
  - 时间序列的唯一标识维度: 维度组合决定一条time series
  - 两条不同的time series: `http_requests_total{method="GET", instance="A"}`与`http_requests_total{method="POST", instance="A"}`

### Label的作用:
  - 支持多维度查询
  - 支持聚合
  - 支持灵活建模

### Label过大的缺点:
  - `高基数问题`: 内存暴涨、查询慢、TSDB压力大; 避免Label的值为唯一值,减少Label

### 高基数:
  - label维度取值过多,导致time series数量暴涨

### Time Series:
  - 唯一性只由`metric name`和`Label set`决定,和时间无关

### Scrape流程:
  - 定时触发Scrape: `Prometheus scheduler`的每个job配置`scrape_interval: 15s`,代表每15秒拉取一次
  - Http请求/metrics: 向目标exporter发起http请求
  - Exporter生成metrics: 例如node_exporter产生的`node_cpu_seconds_total{mode="idle"} 12345`
  - Prometheus解析文本: metric name、label、value、timestamp
  - 转成time series
  - 写入WAL、TSDB: 先写日志、再写缓存、最后写db

### counter是什么:
  - 只增不减
  - 用于累计值(请求数、失败数)
  - 一般不会`自然下降`

### counter reset是什么:
  - 因为counter是进程内存中的累加变量,进程重启或应用重启导致计数器为0
  - Prometheus通过rate自动处理,发现`当前值<上一个值`则直接reset

### 分位数: 
  - 将一组数据从小到大排序,处在某个百分比位置的值
  - P70: 假设一组延迟数据(ms): `10 20 30 40 50 60 70 80 90 100` => P50就是50%的数据小于50ms

### Histogram与Summary区别:
  - Histogram: 服务端预分桶,可聚合,推荐生产使用
  - Summary: 客户端直接算分位数,不可聚合,不推荐大规模使用

### Histogram的数据长什么样:
  - 本质是: `累积分布函数(CDF)`
  - `http_request_duration_seconds_bucket{le=0.1} 100`: 所有小于等于0.1秒的请求数是100条
  - `http_request_duration_seconds_bucket{le=0.3} 300`: 所有小于等于0.3秒的请求数是300条
  - `http_request_duration_seconds_bucket{le=1} 800`: 所有小于等于1秒的请求数是800条

### Histogram为什么必须le:
  - `le`是`Histogram bucket`的`上边界标签`,Histogram需要根据它来确定数据分布区间,否则无法还原分位数
  - 多实例情况下,相同le的bucket必须合并,否则分布会被拆散,无法形成完整的CDF
  - `histogram_quantile`的计算依赖完整的累积分布函数CDF,而`le`是定义分布横轴的唯一维度.如果不按`le`聚合,不同实例的bucket会被错误合并或拆分,导致分布函数不连续,无法正确计算分位数

### sum(rate()) vs rate(sum()):
  - sum(rate()): 先算每条time series的rate,再求和 => 每个series(每个label组合)先算增长速度,再把这些速度相加
  - rate(sum()): 先把所有series合并成一条"大曲线",再整体计算增长速度,已经丢失了label意义

### Recording Rule:
  - 把复杂 PromQL 预计算成新 time series，提高查询性能，降低实时计算成本
  - 由`Rule Manager(规则执行器)`定时调度执行,默认按`evaluation_interval`周期性执行,每个`rule group`独立按周期跑

### PromQL
  - 计算QPS(每秒请求数):  
    - 相关指标:`http_requests_total`
    - 统计最近5分钟的QPS: `rate(http_requests_total[5m])`
  - 按接口统计QPS: 
    - 相关指标: `http_requests_total{path="/api"}`
    - 按接口统计QPS: `sum by (path) (rate(http_requests_total[5m]))`
  - 计算错误率:
    - 相关指标: `http_requests_total{code="200"}` `http_requests_total{code="500"}`
    - 计算错误率: `sum(rate(http_requests_total{code=~"5.."}[5m])) / sum(rate(http_requests_total[5m]))`
  - 找出QPS最高的接口:
    - `topk(5, sum by (path) (rate(http_requests_total[5m])))`
  - 统计过去五分钟新增请求数:
    - `increase(http_requests_total[5m])`
  - 延迟指标:
    - 相关指标: `http_request_duration_seconds_buket`
    - 计算P99延迟: `histogram_quantile(0.99, sum(rate(http_request_duration_seconds_buket[5m])) by (le))`
  - 平均响应时间:
    - `rate(http_request_duration_seconds_sum[5m]) / rate(http_request_duration_seconds_count[5m])`
