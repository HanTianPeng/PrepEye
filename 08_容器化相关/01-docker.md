### Namespace: 一种资源隔离机制，用于把系统资源"划分成多个独立视图“
- 容器本质: 一组运行在同一个内核上的进程，Docker通过`Linux namespace`实现进程、网络、文件系统等资源的隔离，使得容器内的进程拥有独立的运行环境，但本质上仍共享宿主机内核。
- 容器本质: `Linux 进程` + `Namespace` + `CGroups` + `rootfs` + `runtime`

### Namespace类别:
- PID Namespace: 进程隔离
  - 每个容器有独立进程树，容器内PID=1，但在宿主机可能是PID=2345
  
- NET Namespace: 网络隔离
  - 独立IP、端口、路由表，每个容器都有自己的网络栈
  
- MOUNT Namespace: 文件系统隔离
  - 独立文件系统视图，容器看到的是自己的rootfs
  
- IPC Namespace: 进程通信隔离
  - 隔离共享内存、信号量
  
- UTS Namespace: 主机名隔离
  - 每个容器可以有不同的`hostname`

- USER Namespace: 用户隔离
  - 容器内`root`不等于宿主机`root`，更安全

### CGroups: 用来限制、统计、隔离进程资源使用的机制
- `Linux`提供的资源控制机制，用于限制和隔离进程的CPU、内存、IO、网络、进程数等资源使用，Docker通过`cgroups`实现容器级资源配额和限制。
- `CGroups`通过`CFS(Completety Fair Scheduler)`为每个`CGroup`分配`权重`。
- `CFS`通过`虚拟运行时间(vruntime)`来实现进程调度平衡

### Overlay2: 联合文件系统(UnionFS)
- lowerdir(只读层): 镜像的每一层(只读)
- upperdir(可写层): 容器运行时写入的内容
- workdir(工作目录): overlayFS内部使用 
- merged(最终视图): 容器看到的完整文件系统

### 写时复制: Copy-On-Write
- 一个镜像可以被多个容器使用，但是不需要在内存和磁盘上做多个拷贝。
- 在需要对镜像提供的文件进行修改时，该文件会从镜像的文件系统被复制到容器的可写层的文件系统进行修改，而镜像里面的文件不会改变。
- 不同容器对文件的修改相互独立，互不影响。

### 用时分配
- 按需分配空间，而非提前分配，即当一个文件被创建出来后，才会分配空间。

### 为什么删除文件不减少镜像?
- docker镜像是分层存储结构，删除只发生在`upperdir(可写层)`做标记删除，不会较少`lower(只读层)`体积。
- 真正减少镜像体积，需要通过减少layer、合并RUN指令、多阶段构建、使用轻量基础镜像

### docker VS 虚拟机
- 启动、硬盘使用情况、性能、系统支持量、`Guest OS`

### docker run VS docker start
- `docker run`: `docker create` + `docker start` + `docker attach`; 新建 + 启动
- `docker start`: 启动

### docker attach VS docker exec
- `docker attach`: 连接容器的主进程（PID 1）的标准输入输出
- `docker exec`: 在运行中的容器里新启动一个进程

### docker基础命令
- 查看日志
  - 基础命令: `docker logs <container>`
  - 实时日志: `docker logs -f <container>`
  - 最近日志: `docker logs --tail 100 <container>`
  - 带时间戳: `docker logs -t <container>`
- 生产环境进入容器:
  - `docker exec -it <container> bash`
  - `docker exec -it <container> sh`
- 复制文件:
  - 从容器拷贝到宿主机: `docker cp container:/path/file /host/path`
  - 从宿主机拷贝到容器: `docker cp /host/path container:/path/file`
  - 拷贝整个目录: `docker cp container:/data /data` 

### docker run的过程
- 镜像拉取（pull）
- 创建容器（create） 
- 文件系统准备（rootfs / overlay2） 
- namespace 创建（隔离） 
- cgroups 设置（资源限制） 
- 启动进程（exec / PID 1） 
- 运行中

### dockerfile中 CMD VS ENTRYPOINT:
- `CMD`提供容器默认执行的命令或参数，可以被`docker run`覆盖；
- `ENTRYPOINT`定义容器的固定入口程序，是容器启动的主执行逻辑；
- 两者结合使用时`ENTRYPOINT`决定程序，`CMD`提供默认参数；

### dockerfile中 COPY VS ADD:
- `COPY`: 只做文件或目录的简单拷贝
- `ADD`: 在`COPY`的基础上增加了解压tar + URL下载；会降低dockerfile的可预测性，影响构建缓存和可维护性；不推荐使用。 

### docker save VS docker load:
- `docker save`: 把镜像打成一个tar离线包
- `docker load`: 从tar离线包恢复镜像到本地docker

### 多阶段构建: Multi-Stage-Build
- 在一个dockerfile中使用多个`FROM`，把`构建环境`和`运行环境`分离，只把最终产物带入最终镜像。
- `COPY --from=builder /*** /***`
- 解决了哪些问题：
  - 镜像很大，包含编译工具
  - 有安全风险，比如gcc、go等都在
  - 运行环境不干净

### docker layer cache: 是按指令 + 上下文内容做缓存命中
- 在多阶段构建中，每个stage都独立缓存，但可以复用前面的stage缓存结果，一旦某层失效，该层之后的所有层都会失效。
- 利用`.dockerignore`避免无效文件破坏cache。