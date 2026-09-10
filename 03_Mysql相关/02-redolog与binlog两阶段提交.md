#### 01-redolog与binlog的概念
  - `redolog`是innodb存储引擎专属的物理日志，记录的是哪个事务在哪个时间点进行了什么样的修改。核心作用就是保证事务的持久性，用于mysql崩溃之后的快速数据恢复。
  - `binlog`是mysql服务器层面的二进制日志文件，文件有三种格式，分别是statement、row、mixed，核心作用是主从复制和数据备份恢复。

#### 02-那么为什么会出现两阶段提交呢？
  - 其实非常简单，因为两种日志位于不同的层次，当进行对应的SQL操作的时候都会进行日志的记录，那么先写日志，后写磁盘会为问题，因此引入了两阶段提交。
  - 所谓的两阶段提交就是将两个步骤的提交过程拆分为准备阶段和提交阶段两个步骤，根据redolog和binlog的日志顺序，要么两者都完整的执行到磁盘中，要么两者都不执行，不会出现一个成功，一个失败的情况，这种机制是为了保证数据一致性和可用的。

#### 03-redolog和binlog的两阶段提交
  - 当开启事务之后，执行SQL语句，会记录对应的redolog，此时redolog位于`redo log buffer`，持久化要提交的数据，然后修改`buffer pool`中的数据，然后再记录binlog，同样，binlog的持久化也是由参数配置的。 
    - 执行commit指令之后，innodb存储引擎会将当前事务的`redo log buffer`中的redolog进行持久化操作，此时redolog中记录了数据修改相关的信息，此时的日志状态为prepare，而不是commit，这步完成之后，binlog日志也开始进行对应的持久化操作。 
  - 当mysql的server将当前事务的binlog持久化成功后，会向innodb存储引擎发送提交确认的消息，innodb收到之后，会将之前的prepare状态的redolog文件修改为commit状态。 
  - innodb可以释放对应的资源，向客户端返回事务执行成功的响应，整个事务提交流程结束。
  - `redolog`的刷盘配置:`innodb_flush_log_at_trx_commit`
  - `binlog`的刷盘配置: `sync_binlog`
  - 两阶段提交保证的是`redo log`和`binlog`的逻辑一致性顺序，而不是强制的物理刷盘顺序；真正的持久性取决于`innodb_flush_log_at_trx_commit`和`sync_binlog`的配置。

#### 03-文件
  - 逻辑层(file): 有名字(ib_logfile0)，操作系统认识，可以open、close，但不是数据本身。
  - 内存层(os page cache): 内核态文件系统缓存
  - 物理层(disk): 磁盘

#### 04-三个核心组件:
  - buffer pool: 数据页，innodb的内存数据页缓存，真正的数据修改发生在这里(dirty page)
  - redolog buffer: 记录`我改了那些数据页`，物理日志，保证crash后可以恢复buffer pool；当buffer pool修改后开始写入redolog buffer
  - binlog cache: 记录`SQL做了什么`，逻辑日志，作用:主从复制；

#### 03-完整的提交链路
  - buffer pool修改数据页(dirty page)
  - 写入redolog buffer
  - prepare redolog
  - write binlog
  - sync_binlog 决定是否fsync
  - commit redolog

#### 04-redo刷盘行为
  - 用户态内存: `redolog buffer`
  - 内核态文件系统缓存: `write(写入OS Cache)`
  - 磁盘: `fsync(真正落盘)，产生 redolog file`

#### 05-redolog的刷盘配置innodb_flush_log_at_trx_commit:
  - 0: 性能最高，最危险；每秒进行一次 redolog buffer -> fsync；commit时，不进行刷盘。
  - 1: 每次commit，redolog buffer -> os cache -> fsync。
  - 2: 每次commit，redolog buffer -> os cache；后台每秒fsync。

#### 06-binlog刷盘行为
  - 内存: `binlog cache`
  - 文件系统缓存: `write(写入OS Cache)`
  - 磁盘: `fsync(真正落盘)，产生binlog file`

#### 07-binlog的刷盘配置sync_binlog:
  - 0: 由操作系统决定什么时候刷盘，可能丢失多个事务的binlog
  - 1: 每次commit，binlog cache -> os cache -> fsync。
  - N: 每N次commit，binlog cache -> os cache -> fsync，可能丢失N个事务的binlog。

#### 06-WAL(Write-Ahead Logging)
  - 是一种日志优先机制：在修改数据页之前必须先写`redo log`，通过日志保证即使发生crash，也可以通过重放日志恢复数据库状态，从而实现高性能与高可靠性的平衡。

#### 04-至于为什么要给日志添加状态来保证数据一致性，可以通过反证法的方式验证。
  - 先写binlog，然后再写redolog
    - 假设binlog已经成功刷盘，但是redolog还没有来得及刷盘，此时mysql突然崩溃了，重启之后，innodb存储引擎因为没有对应的redolog记录，会将事务回滚，数据保留一份，但是binlog中已经有了这条事务的执行记录，在重做上述操作的过程中，从库就会丢失一条记录，主从的数据不一致了。
  - 先写redolog，再写binlog
    - 假设redolog已经成功刷盘，binlog还没有刷盘，mysql崩溃之后，重启，innodb引擎通过redolog进行数据恢复的时候，该事务会被提交，数据恢复后，但是binlog中没有这条记录，主从复制的时候也不会同步这条数据，那么就会出现主库有，从库没有的情况，主从数据不一致了。

#### 05-通过对redolog添加状态的方式就可以解决这个问题，如何对崩溃的恢复呢？
  - 当mysql崩溃重启之后，innodb引擎会自动触发崩溃恢复的逻辑，核心是扫描redolog，根据redolog的状态对binlog的完整性、一致性要求高的部分进行恢复，从而确保数据的一致性。
    - redolog中没有记录，或者只有部分记录，没有prepare标记，说明事务还没有开始，直接回滚即可，没有数据一致性的问题。 
    - redolog的状态是prepare，但是找不到对应的完整binlog，说明redo阶段完成，但是提交阶段没有完成，此时会回滚事务，保证redolog和binlog数据一致。 
    - redolog为prepare状态，能够找到与之匹配的binlog日志，说明binlog已经刷盘，只是还没有来得及修改redolog的状态为commit，此时innodb会自动完成commit操作，恢复事务，确保数据一致。

