#### 大Key问题
  - 内存浪费及数据倾斜
  - 网络带宽阻塞
  - 慢查询与性能雪崩
  - 高危删除操作

#### 解决方案
  - 拆分: 将一个大的、复杂的对象拆分成多个小的、独立的Key，降低单Key的内存占用。
  - 压缩: 对无法拆分的大Value进行序列化压缩(Gzip、Snappy)，显著减少网络传输与内存占用，由于需要消耗CPU进行压缩和解压，所以应用场景针对于访问频率低，
  - 选型: 根据业务场景，选择Hash、List等更合适的Redis数据结构，替代单一的String类型。
  - 清理: 建立定期巡检与监控机制，从源头识别，减少和清理无效的大Key。删除采用`UNLINK`，主线程断开引用，交给`lazyfree线程`，不阻塞主线程执行命令。

#### LazyFree配置:
  - `lazyfree-lazy-expire yes`: 过期删除
  - `lazyfree-lazy-eviction yes`: 淘汰删除
  - `lazyfree-lazy-server-del yes`: DEL删除