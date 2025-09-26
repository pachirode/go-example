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


