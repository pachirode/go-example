# flag

命令行参数解析标准库

### 使用

##### flag.Type()

- 标志名称
- 标志默认参数
- 标志使用帮助信息

##### flag.TypeVar()

取消返回值，将用户传递的命令行参数绑定到第一个参数

- 需要绑定的参数变量
- 标志名称
- 标志帮助信息

##### flag.Var()

可以自定义类型，需要实现 `flag.Value`

- 被绑定的命令行参数
- 标志名称
- 标志帮助信息

##### 支持标志类型

- `bool`
- `time.Duration`
- `int`
- `string`
- `flag.Value`

### 源码

##### 定义标识

flag.Int

底层代理调用 `CommandLine` 方法

CommandLine

`FlagSet` 结构体指针，构造器需要传入两个参数

- `os.Args[0]`
    - 当前命令行程序的名称
- `ExitOnError`
    - 常量
    - 标记出现 `error` 做法，表示遇到 `error` 时退出程序

##### FlagSet

- `Usage func()`
    - 指向 `--help/-h` 参数，查看命令行程序使用帮助时会被调用
- `parsed bool`
    - 标记是否调用过，`flag.Parse()`
- `actual map[string]*Flag`
    - 命令行解析的标志参数
- `formal map[string]*Flag`
    - 程序默认指定的标志参数
- `args []string`
    - 保存处理完标志剩余的参数列表
- `errorHandling ErrorHandling`
    - 标记出现 `error` 如何处理
- `output io.Writer`
    - 设置输出位置

##### Flag

用来记录一个命令行参数，存储一个标志的所有信息

- `Name string`
- `Usage string`
- `Value Value`
- `DefValue string`

##### Value 接口

使用接口是为了可以存储任意类型的值，支持用户自定义类型

- `String() string`
- `Set(string) error`

##### 解析标志参数

底层调用的 `CommandLine.Parse`

- 首先将 `f.parsed` 标记为 `true`
- 将参数保存到 `f.args`
- 循环解析命令行参数，每次解析一个标志直至解析完成或者遇到错误

##### parseOne

- 对 `f.args` 参数进行校验
- 提取标志前符号 `-` 的个数放到 `numMinuses` 变量
- 取出标志，对标志做语法检查
- 取出参数值，并判断参数名是否为 `--help/-h`
    - 是，打印使用信息并退出
- 根据参数值是否为 `bool` 进行参数绑定，将参数设置到对应的标志变量中，并将标志保存到 `f.actual`


# pflag

作为标准库的替换，兼容标准库，可以直接使用 `flag` 包定义的全局对象 `CommandLine`

### 使用

基本用法和 `flag` 保持一致

### 额外用法

##### 支持设置简短标志
- `--`
  - 完整标志
- `-`
  - 简短标识符

##### 标志名 Normalize

可以给标志起一个或者多个别名

##### NoOptDefVal

创建标志之后可以设置无默认值

##### 弃用/隐藏标志 

