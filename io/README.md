# Discard

用于丢弃所有写入的数据

### 丢弃响应体内容

操作失败或者未读取 `Body`，`Transport` 不会复用 `TCP` 连接

### 源码

`Discard` 是一个变量，实现了 `io.Writer` 接口，实现类似 `unix` 上 `/dev/null` 的功能