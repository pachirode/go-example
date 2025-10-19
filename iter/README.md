# 实现迭代器

### 迭代器模式

### 回调函数

### 通道

### 官方迭代器

`for-range` 支持的

- `array`
- `slice`
- `string`
- `map`
- `channel`
- `integer`
- 特定函数
    - `func(func() bool)`
    - `func(func(V) bool)`
    - `func(func(K, V) bool)`
- `iter`
  - 官方迭代器库

##### 使用泛型

实现泛型函数就能迭代同数据结构的所有对象

### Push & Pull 迭代器

##### Push

由于迭代器自身控制迭代进度，迭代器负责迭代逻辑，主动将元素推送给 `yield`

##### Pull

通过 `next` 的方式，由调用方主动获取元素

实现涉及到 `coro` 空结构体，使用 `linkname` 从 `runtime` 中获取，实际上是一种更加轻量级的协程，提供了 `coroswitch` 可以主动切换这种协程
这种模式同时只会有一个协程在运行

- 创建一个 `coroutine`，主协程运行 `next`，`stop` 函数用来终止迭代
- 主协程不会立即运行，需要调用显式的调用 `next` 恢复
- 协程开始运行，之后反转让出执行权限给 `next`