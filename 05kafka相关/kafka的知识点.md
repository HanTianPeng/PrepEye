## 消息中间件三大功能
- 异步
- 解耦
- 流量削峰

## 四大核心组件
- Producer: 生产者,向Kafka broker发消息的客户端
- Consumer: 消费者,向Kafka broker取消息的客户端
- Broker: Kafka的服务节点,存储信息;一台Kafka服务器就是一个broker,一个集群由多个broker组成。一个broker可以容纳多个topic。
- ZooKeeper/KRaft(Kafka Raft): 管理集群元数据(Kafka3.3+已支持去除ZooKeeper)
- Kafka正在移除对ZooKeeper的依赖,使用内置的KRaft模式(基于Raft协议)管理元数据

## Kafka架构
- Topic: 可以理解为一个队列,一个Topic又分为一个或多个分区

## Kafka为什么这么快
- Kafka是分布式消息系统,需要处理海量的消息，Kafka的设计是把所有的消息都写入速度低容量大的硬盘,以此来换取更强的存储能力.
- 但实际上使用硬盘并没有带来过多的性能损失。主要是因为以下几点:
  - 磁盘顺序读写
  - 零拷贝(Zero-Copy): 通过sendFile()系统调用,数据直接从磁盘->网卡,不经过用户态内存;
  - 文件分段
  - 批量压缩+批量发送: Producer批量攒消息,减少网络请求;支持gzip/snappy/lz4压缩;

## Kafka的缺点
- 由于是批量发送,数据并非真正的实时
- 对于mqtt协议并不支持
- 不支持物联网传感数据直接接入
- 仅支持统一分区内的消息有序,无法实现全局消息有序
- 监控不完善,需要安装插件
- 依赖ZooKeeper进行元数据管理

## Partition(分区)的作用是什么
- 为了实现扩展性,一个非常大的Topic可以分布到多个Broker上,每个Partition是一个有序的队列;
- Broker的数量决定了最大副本数,不是Partition数量;Partition数量可以远远大于Broker的数量,但Replication数量不能超过Broker数量
- Partition中的每条消息都会被分配一个有序的id(offset);
- 将消息发给Consumer,Kafka只保证一个Partition中的消息的顺序,不保证一个Topic的整体(多个Partition间)的顺序;
- 水平扩展: 一个Topic可拆分成多个Partition,分布到不同的Broker;
- 并行消费: 一个Consumer Group中,每个Partition只能被一个Consumer消费;
- 保证局部有序: 同一个Partition内消息有序(但全局无序);

## Kafka的分区数是否可以增加或减少
- 可以使用bin/kafka-topic.sh命令对Kafka增加Kafka的Topic分区数,但是不支持减少
- 减少的分区其数据如何处理?删除就会导致消息丢失，保留如果追加到其他的分区就破坏了Partition的有序性

## Kafka能否保证全局有序
- 不能: 只能保证单Partition有序
- 如果业务需要全局有序:
  - Topic只设1个Partition(牺牲并发)
  - 或则在消息中加入业务ID,由消费者合并排序(复杂,冗余业务层)

## 消息在磁盘上怎么存储的
- 存储结构
  - 每个Partition对应一个目录(如topic-0/)
  - 目录下包含多个Segment文件(.log数据文件 + .index偏移索引 + .timeindex时间索引)
  - Segment按大小(默认1GB)或时间滚动
- 查找消息流程
  - 根据offset定位到Segment
  - 用.index二分查找物理位置
  - 从.log读取数据
- 优点: 避免单文件过大,便于清理过期数据(直接删除旧的Segment文件)

## Producer发送消息有几种模式(数据传输的事务定义)
- 最多一次: acks=0 可能丢，吞吐最高;消息不会被重复发送,最多被传输一次,但也可能一次不传输
- 至少一次: ack=1, 不丢但可能重复, 吞吐高;消息不会被漏发送,最少被传输一次,但也有可能被重复传输
- 精准一次: ack=all+幂等/事务, Exactly-Once, 吞吐中;不会漏传输也不会被重复传输
- 生产建议: 日志类(ack=1); 订单类(ack=all+幂等)

## 如何保证消息不丢失
- Producer: ack=all, retries=Integer.MAX_VALUE, enable.idempotence=true
  - maxInFlightRequestsPerConnection: 已经发出去但还没有收到响应的Request对象
- Broker: replication.factor>=3, min.insync.replicas=2
- Consumer: 手动提交offset(处理完再commit)
- 错误观念: 只要Producer发送成功就不会丢,因为如果Broker只有一个副本且宕机,照样丢

## 什么是幂等Producer,如何实现
- 幂等性: 即使重试多次,Broker也只会保留一条消息
- 实现原理:
  - 每个Producer有唯一的PID(重启会变)
  - 每条消息会带sequence number (序列号)
  - Broker为每个<PID Partition>维护序列号,拒绝重复或乱序消息
  - 幂等性仅保证单Partition会话内Exactly-Once,跨Partition或重启无效

## Consumer Group是什么,有什么作用
- 多个Consumer组成一个Group,共同消费一个Topic
- Kafka保证每个Partition只分配给Group内一个Consumer(也就是说同一个消费者组内,消费并行度由Partition数决定)
- 如果Partition少于Consumer数量,对于的Consumer会处于idle状态(空闲)
- 如果Partition多余Consumer数量,每个Consumer会分到多个Partition,所有的Consumer均工作
- 实现负载均衡+容错(Consumer挂了,Partition会reassign)
- 应用场景: 
  - 广播: 每个服务实例独立消费 -> 不同的Group
  - 负载均衡: 多个实例分摊消费 -> 同一个Group

## offset是什么
- Kafka的存储文件都是按照offset.kafka来命名的,用offset做名字的好处是方便查找.
- 例如你想找2049的位置,只需要找2048.kafka的文件即可
- the first offset就是00000000000.kafka

## offset提交方式有哪些,如何避免重复/丢失
- 自动提交: enable.auto.commit=true, 定时提交(如每5秒), 处理中如果宕机,重复消费
- 手动提交: commitSync/commitAsync, 业务处理完再提交, 忘记提交导致重复消费, 提交太早导致丢失

## Kafka如何实现Exactly-Once语义
- Producer: ack=all, retries=Integer.MAX_VALUE, enable.idempotence=true
- 使用Kafka Streams或Transactional Producer: 通过事务保证读->处理->写原子性

## ISR、OSR、AR是什么
- ISR: In-Sync Replicas 副本同步队列
  - 由Leader维护,Follower从Leader同步数据有一些延迟,超过相应的阈值会把Follower剔除出ISR,存入OSR列表;
  - 新加入的Follower也会先存放在OSR中;AR=ISR + OSR
  - replica.lag.time.max.ms: 延迟时间,副本在复制消息时可被延迟的最长时间
  - replica.lag.max.messages: 副本延迟的消息最大个数;由于需要评估Producer写入效率,而且还存在流量高峰期场景;这个参数实际意义不大
- OSR: Out-of-Sync Replicas
- AR: Assigned Replicas 所有副本

## LEO、HW、LSO、LW等分别代表什么
- LEO: LogEndOffset的简称,代表当前日志文件中的下一条
- HW: High Watermark的简称,取Partition对应的ISR中最小的LEO作为HW,Consumer最多只能消费到HW所在的位置上一条消息
- LSO: LastStableOffset的简称,对于未完成的事务而言,LSO的值等于事务中第一条消息的位置(FirstUnstableOffset);对于已完成的事务而言,等同于HW
- LW: Low Watermark的简称,代表AR集合中最小的LogStartOffset值

## Kafka数据一致性
- HW

## Kafka消费者是否可以消费指定分区消息
- Kafka Consumer消费消息时,向Broker发出fetch请求去消费特定分区的消息,Consumer指定消息在日志中的偏移量(offset),就可以消费从这个位置开始的消息
- Consumer拥有offset的控制权,就可以向后回滚消费之前的消息

## Kafka创建Topic时如何将分区放置到不同的Broker中
- 副本因子不能大于Broker的数量(Broker、Partition、Replicas 三个概念必须要整明白)
- 第一个分区(编号为0)的第一个副本放置位置是随机从BrokerList选择的,startIndex
- 其他分区的第一个副本放置位置相对于第0个分区依次往后移。
- 也就是说如果我们有5个Broker,5个分区,假设第0号分区放在第四个Broker上，第1号分区就放在第五个Broker上，第2号分区就放在第一个Broker上,以此类推。
- 剩余的副本相对于第一个副本放置位置其实是有nextReplicaShift决定的,而这个数是随机产生的;
- https://blog.51cto.com/u_15127589/2680649
- assignReplicasToBrokers(Seq(0, 1, 2, 3), 3, 3),这个方法的第一个参数:BrokerList,第二个参数:Partition,第三个参数:Replicas
- 分别产生3个分区:
- (0, List(3, 2, 0))  编号为0的分区对应的三个副本分别存在3号Broker、2号Broker、0号Broker
- (1, List(0, 3, 1))
- (2, List(1, 0, 2))
- assignReplicasToBrokers(Seq(0, 1, 2, 3), 5, 3)
- 分别产生5个分区:
- (0, List(3, 1, 2))
- (1, List(0, 2, 3))
- (2, List(1, 3, 0))
- (3, List(2, 0, 1))
- (4, List(3, 1, 2))

## Kafka的副本复制机制
- Kafka主题中的每个分区都有一个预写日志(write-ahead log),我们写入Kafka的消息就存储在这里面。这里面每一条消息都有一个唯一的偏移量，用于标识它在当前分区日志中的位置。
- Kafka主题中的每个分区都被复制了n次，其中n就是主题复制因子(replication factor)。这允许Kafka在集群服务器发生故障时自动切换到这些副本，以便在出现故障时消息仍然可用。
- Kafka的复制是以分区为粒度的，分区的预写日志被复制到n个服务器上。
- 在n个副本中，一个副本作为Leader，其他副本成为Followers。
- Producer只能往Leader分区写数据(读也只能从Leader分区进行)，Followers只按顺序从Leader上复制日志。
- 每个分区的Leader会维护一个ISR(In-Sync replicas 同步副本列表)，当Producer往Broker发送消息时，消息先写入到对应Leader分区上，然后复制到这个分区的所有副本中。
- 只有将消息成功复制到所有同步副本(ISR)后，这条消息才算被提交(Commited / LogEndOffset)。
- 由于消息复制延迟受到最慢同步副本的限制，因此快速检测慢副本并将其从ISR中删除非常重要。
- 一个副本如果它没有跟上Leader的日志进度(Caught Up)，那么它可能会被标记为不同步的副本。
- https://blog.51cto.com/u_15127589/2677435

## 什么情况会导致一个副本与Leader失去同步
- 慢副本(Slow Replica): Follower Replica在一段时间内一直无法赶上Leader的写入速度。造成这种情况的最常见原因之一是Follower Replica上的I/O瓶颈，导致它持久化日志的时间比它从Leader消费消息的时间长。
- 卡住副本(Stuck Replica): Follower Replica在很长一段时间内停止从Leader获取消息。这可能是GC停顿或者副本发生故障。
- 刚启动副本(Bootstrapping Replica): 新增加的Replica是不同步的，直到跟上Leader的日志。

## Kafka的再均衡
- 在Kafka中,当有新消费者加入或者订阅的Topic数发生变化时, 会触发Rebalance(再均衡:在同一个消费者组中,分区的所有权从一个消费者转移到另外一个消费者)机制;
- Rebalance:重新均衡消费者消费。Rebalance的过程如下:
  - 所有成员都向Coordinator发送请求,请求入组;一旦所有成员都发送了请求,Coordinator会从中选择一个Consumer担任Leader角色,并把组成员信息及订阅信息发给Leader;
  - Leader开始分配消费方案,指明具体哪个Consumer负责消费哪些Topic的哪些Partition。一旦分配完成.Leader会将这个方案发给Coordinator。
  - Coordinator接收到分配方案之后会把方案发给各个Consumer,这样组内的所有成员都知道自己消费哪些分区了.