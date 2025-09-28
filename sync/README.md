# errgroup

用于并发的执行多个 `goroutine`，基于 `sync.WaitGroup` 基础实现

### 对比 `sync.WaitGroup`

##### 错误处理

`sync.WaitGroup` 只负责等待 `goroutine` 完成，不处理返回值或者错误

`errgroup.Group` 不能直接处理 `goroutine` 的返回值 ，但是一个返回错误时，可以取消其他 `goroutine`，并返回第一个非 `nil` 的错误

##### 上下文取消

`errgroup` 可以和 `context.Context` 配合使用，在某个 `goroutine` 失败时自动取消其他的

##### 简化并发编程

`errgroup` 可以减少错误处理的样板代码，开发者不需要手动管理错误状态和同步逻辑

##### 限制并发数量

提供便捷接口来限制并发 `goroutine` 数量，`SetLimit`

##### 尝试启动

`TryGo` 尝试启动一个任务，返回一个 `bool` 值，识别任务是否成功启动
需要配合 `SetLimit`，如果不限制并发数量，会始终返回 `true`

# sync.Once

一种同步原语，确保一个操作在程序的生命周期只能被执行一次

`Once` 对象不能被复制

在多个 `goroutine` 中，`once.Do(f)` 首次被调用，函数会执行，并保证它的执行结果在内存中对其他的 `goroutine` 都是可见的
之后再次调用这个方法，不会重新执行函数，它们会同步函数的结果

### 使用场景

- 单例模式
- 懒加载
- 并发安全初始化

### 源码

结构体

- `_ noCopy`
    - 标识这个结构体不可以复制
- `done atomic.Uint32`
    - 故意被放在结构体第一个字段，可以减少 `CPU` 执行的指令数，优化性能
- `m    Mutex`

`Do` 方法

只有第一执行才会进入到 `slow-path`，将 `slow-path` 分离出来是为了对 `fast-path` 进行内联优化

- `fast-path`
    - 针对常见操作或最佳情况进行优化的代码路径
    - 需要避免高成本操作
        - 加锁
        - `IO` 操作
- `slow-path`
    - 处理罕见复杂的情况，通常执行步骤较多

# sync.OnceFunc

返回一个仅调用 `f` 一次的函数，返回的函数可以并发调用
如果函数调用出现 `panic`，返回的函数每次调用都会产生相同的 `panic`

# sync.OnceValue

使用了泛型，接收函数带有返回值，返回函数可带有返回值

# sync.WaitGroup

阻塞等待一组并发任务的完成，会在内部维护一个计数器，`wg.Add(1)` 计数器加一，`wg.Done` 计数器减一，`wg.Wait()`
阻塞调用者 `goroutine`，直到计数器为 0
`sync.WaitGroup` 零值可用

### 源码

- `nocpoy`
    - 用于标识结构体不能被复制
- `state`
    - 原子类型，高 32 位是计数器的值，低 32 位是等待者的数量
- `sema`
    - 信号量，阻塞和唤醒 `waiter`

### Done

计数器 `counter` 减 1，相当于调用 `wg.Add(-1)`

### Add

通过移位操作获取 `counter` 和 `waiter` 的值
校验 `counter` 不能为负，`Add` 和 `Wait` 不能并发调用，否则会触发 `panic`
计数器为 `0`，唤起 `wg.Wait` 阻塞，在此之前会再次进行校验，如果添加任务会触发 `panic`

### Wait

使用无限循环，重试 `CAS` 操作及时 `CompareAndSwap`，如果成功 `waiter` 数量加一

# sync.Cond

并发原语，通过一个条件来实现阻塞和唤醒一组需要协作的 `goroutine`
调用 `Wait` 时，当前 `goroutine` 会被阻塞，直到其他调用 `Broadcast` 或者 `Signal`

### 源码

- `noCopy`
  - 静态检查
- `L`
  - 互斥锁，修改条件时需要持有
- `notify`
  - 记录被阻塞的等待队列，维护一个通知列表
  - 协调 `goroutine` 的阻塞和唤醒
- `checker`
  - 同样防止结构体被复制，运行时进行动态检查

# 内联优化

编译时的优化技术，它是将被调用的函数体直接嵌入到调用处，而不是生成一次真正的函数调用，这样可以减少函数调用的开销

`go build -gcflags='-m'` 构建参数可以查看内联情况，日志中 `inlining` 表示使用了内联优化

# singleflight

官方扩展库提供的扩展并发的原语，可以将多个并发请求合并为一个，主要作用是抑制重复的并发调用

如果多个 `goroutine` 并发调用同一个函数，`singleflight`
可以只让一个发起调用，其他的都阻塞住，等调用结果返回，在同时返回给多个 `goroutine`

读操作较为常见，写操作需要慎重

### 应用场景

大量的请求读取 `redis` 缓存，发现缓存失效，继续向下请求 `mysql`
此时可以使用 `singleflight` 合并请求，只保留一个请求去调用 `mysql` 数据库，然后将结果返回给所有请求

### 源码

##### Group

- `mu sync.Mutex`
    - 互斥锁，用来保护下面 `m` 的访问
- `m  map[string]*call`
    - 键是调用 `singleflight.Do` 传入的第一个参数 `key`

方法

- `Do`
- `DoChan`
- `Forget`
    - 忘记一个 `key`，再次调用上面两个方法，不会等待之前未完成的函数执行结果

##### call

- `wg sync.WaitGroup`
- `val interface{}`
    - 记录函数返回值
- `err error`
    - 记录函数错误
- `dups  int`
    - 从缓存中获取需要返回的次数
- `chans []chan<- Result`
    - 为 `DoChan` 提供返回值

一个正在执行的 `in-flight` 或者已经完的 `fn` 函数的调用

##### doCall

- 双层 `defer` 设计
    - 第一层用于捕获 `panic`
    - 第二层用于处理 `runtime.Goexit` 和资源的释放
- `panic` 处理
    - 通过 `goroutine` 执行 `panic(e)` 保证不会阻塞 `channel` 调用

##### 使用场景

缓存击穿

缓存中某个热点键过期，导致大量请求同时访问数据库，导致系统高压
确保缓存重建过程中，只有一个请求会访问数据库

远程服务调用
多个并发请求访问同一个远程服务
使相同的请求合并为一次

定时任务去重
分布式系统中，可能多个节点同时执行定时任务
只有一个节点执行任务，其他节点共享结果

消息去重
消息队列中存在重复消息的消费问题
消费端保证相同的消息只会处理一次

分布式锁优化
多个节点同时抢锁，可能会引发大量重试的加锁操作
降低对分布式锁的访问压力，只允许一个请求进行加锁操作
