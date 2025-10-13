# log

标准日志库

### 输出日志

提供九种方法来输出日志

### 定制日志

`logger := log.New(os.Stdout, "[Debug] - ", xxx)`

- 日志输出位置
    - `io.Writer` 对象
- 日志前缀
- 日志属性
    - 但是预定义好的常量，不能修改，可以通过 `|` 运算符来组合使用

### 源码

基本都是一致的，都是通过 `std.Output` 输出日志
后续的处理结果不同

- `PrintX` 输出日志之后就结束
- `FatalX` 输出日志之后会执行 `os.Exit(1)`
- `Panic` 输出日志之后会执行 `panic(s)`

`std` 其实就是使用 `New` 创建一个 `Logger` 对象，日志输出到标准错误输出，日志前缀为空，属性 `LstdFlags`

```text
//基本格式
func X(V... any) {
    std.Output(2, fmt.Sprint(v...))
}
```

##### Logger

- `mu`
- `prefix`
- `flag`
    - 记录日志属性
- `out`
- `buf`
- `isDiscard`
    - 用来丢弃日志

> `io.Discard` 对象实现了 `io.Writer`, 其执行 `Write` 操作之后不会产生任何实际的效果，是一个用于丢弃数据的对象

##### Output()

- 根据日志属性来决定是否需要获取文件名和行号
    - 调用 `runtime.Caller` 是一个耗时的操作，为了增加并发性，暂时释放锁
    - 获取到文件和行号再重新加锁
- 清空 `buf` 中保留的上次日志信息，通过 `formatHeader` 方法格式化日志头信息
- 将日志内容追加到 `buf`
- 保证日志以 `\n` 结尾
- 通过 `l.out.Write` 将 `buf` 内容输出

##### formatHeader()

按位与来计算是否设置了某个日志的属性，根据日志属性来格式化头信息

# slog

结构化日志库，主要解决 `log` 日志不是序列化和不支持日志级别

`slog` 日志默认输出到 `os.Stdout`，默认级别为 `Info`

### 附加属性

`slog` 支持在 `msg` 后面传入无限多个 `key/value` 键值对来附加额外的属性

### 结构化日志

- `JSONHandler`
    - 支持 `json` 结构化日志
    - 生产环境使用，方便采集处理
- `TextHandler`
    - 将日志输出为 `key=value` 结构
    - 测试环境使用，方便查看

### 自定义 `Logger` 替换默认

可以使用自定义的 `logger` 来替换 `slog` 默认的 `logger` 对象，设置默认对象之后 `log.Logger` 也会被影响

### 使用宽松类型可能出现不匹配的属性键值对

如果后面传入的 `key/value` 附加属性不是成对的，不会报错，但是日志里面会提示
使用强类型可以避免问题

### 子 logger

使用 `with` 方法附加自定义属性到新的 `*slog.Logger` 对象，所有的日志记录都会携带附加的属性
也可以使用属性分组

### 实现 slog.LogValuer 接口，隐藏敏感信息

对象只要实现这个接口，可以直接传递给 `slog` 进行记录

### 源码

##### Logger

提供日志记录的方法

- `Handler` 接口

##### Record

一个实例代表一条日志记录

##### Handler 接口

处理 `Logger` 产生的日志

##### 使用

- 调用 `Logger` 提供的日志记录方法 `Info`
- `Info` 内部调用一个私有的 `log`
- `log` 创建一个日志条目 `Record`
- `Logger` 调用嵌入的 `Handler` 方法解析并执行日志记录逻辑

# zap

轻量级日志记录库，内部避免使用 `interface{}` 和反射来提高代码效率

特点

- 快
- 支持结构化日志记录
- 支持其中日志级别
- 支持输出调用堆栈
- 支持 `Hooks` 机制

### 基本使用

针对生产环境和开发环境提供不同的函数来创建 `Logger` 对象
如果需要附加属性值，需要采用采用强类型，底层使用 `zapcore.Field` 类型来记录

### SugaredLogger

提供更为人性化接口，日志中追加 `key-value` 时不需要使用强类型指定，只需要保证 `key` 是 `string`

### 定制 Logger

只需要定制 `zap.Config`

### 源码

##### 实现 NewProduction() 和 NewDevelopment()

先创建一个配置对象 `zap.Config`，然后再调用配置对象的 `Build` 方法构建 `Logger`

# logrus

特点
- 兼容标准库 `API`
- 可扩展机制 `Hooks` 机制
  - 将不同的日志添加 `Hooks` 记录到不同位置
    - `Error` 发送给特定的程序进行报警
- 并发安全
- 控制台可以输出不同颜色

### Hooks

##### Hook 接口

- `Levels() []Level`
  - 如果记录的日志级别存在，触发 `Hooks` 调用 `Fire` 方法
- `Fire(*Entry) error`

实现这两个接口可以自定义 `Hook`，有大量内置和第三方的
