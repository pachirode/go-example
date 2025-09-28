# Context

主要涉及用来进行链路控制和安全传值，其是并发安全的

### Context 接口

最基本接口，包中提供了多种包的实现

##### emptyCtx

一个空的实现，没有链路控制能力，也无法安全传值，可以作为其他实现的基类

##### cancelCtx

带取消功能的实现

##### withoutCancelCtx

类似 `emptyCtx`

##### valueCtx

用来安全传值

### 方法

##### Background()

构造一个 `backgroundCtx` 对象并返回，通常作为 `context` 树的根节点

##### TODO()

构造一个 `todoCtx` 对象，同样作为 `context` 树的根节点，不知道该用什么的时候，可以用它

##### WithCancel() 和 WithCancelCause()

构造 `cancelCtx`，区别在于是否传入根因

##### WithDeadline() 和 WithDeadlineCause()

用来构造返回 `cancelCtx` 或者 `timerCtx` 可以接收一个 `time.Time` 指定 `context` 的取消时间

##### WithTimeout() 和 WithTimeoutCause()

接收一个 `time.Duration` 来指定多少时间之后 `context` 对象会被取消
内部分别调用上面两个方法

##### WithoutCancel()

##### WithValue()

##### AfterFunc()

在 `context` 过期时异步执行一个任务，会构造一个 `afterFuncCtx` 但是不返回，而是返回一个停止函数，可以阻止异步任务执行

# 源码

### Context 接口

- `Deadline()`
    - 返回该 `context` 应该被取消的截止时间，如果没有设置截止时间，则返回 `ok` 为 `false`
- `Done()`
    - 返回一个只读的取消信号，被取消 `channel` 会被关掉
- `Err()`
    - 被取消原因，如果未被关掉，返回 `nil`
- `Value(key any)`
    - 返回给定键的关联值

### cancelCtx

- `Context`
- `mu`
    - 互斥锁，保证安全的操作结构体
- `done`
    - 为了支持原子操作，减少互斥锁的使用频率
- `children`
    - 集合，记录所有子上下文，父上下文取消的时候，所有的子上下文也可以取消
- `err`
    - 上下文取消的原因
- `cause`
    - 上下文取消的根因

内嵌了 `Context` 接口，支持任意的其他类型的 `context` 实现作为父上下文

### canceler interface

- `cancel`
    - 取消函数
- `Done`
    - 通过返回 `channel` 能知道是否被取消

该接口表示一个可以被取消的对象，支持取消的 `context` 都需要提供这两个方法
父上下文调用取消会调用子上下文的 `cancel` 方法进行级联取消，必须实现 `Done` 才能直到是否被取消完成