# 依赖注入

依赖注入是一种软件的设计模式，允许将组件的依赖项外部化，使得组件本身更加可测
依赖注入可以解耦代码

# wire

自动依赖注入框架，通过代码生成而不是运行时的反射来实现
注入的代码在编译时以及确定

### 导入依赖

- `github.com/google/wire`
- `github.com/google/wire/cmd/wire`
    - 命令行工具，用于生成依赖注入代码

### 使用 `wire`

函数内部有明显的依赖关系，`NewEvent` 依赖 `NewGreet` 依赖 `NewMessage`，将这部分代码进行抽离，封装到 `InintEvent`

一般 `wire` 部分的代码存放在 `wire.go` 文件中，不需要手动的调用构造函数，将他们依次传递给 `wire.Build`

### 生成代码

`wire gen .`

- 扫描指定包使用 `wire.Build` 代码，自动生成 `wire_gen.go`，生成的代码自动完成依赖注入代码，我们只需要将构造函数传入

### 构建约束

`//go:build`，必须放在文件最开始的注释代码，提示 `go build` 如何构建代码
`wireinject` 是传递给构建约束的选项，相当于一个 `if`，可以根据选项来定制构建如何处理 `Go`

此处

- 将文件标记为 `wire` 处理的目标
    - 告诉 `wire` 该文件包含使用了 `wire` 进行依赖注入
    - 文件通常包含 `wire.Build` 函数调用
- 条件编译
    - 保证在正常的构建中该文件最终不会被编译进可执行文件

### 生成的文件

- `//go:generate go run -mod=mod github.com/google/wire/cmd/wire`
    - 用于在编译前自动执行生成代码的命令
- `//go:build !wireinject`
    - 和前面的约束关键取反，告诉 `wire` 忽略这个文件

### 核心概念

- `providers`
    - 提供者
    - 必须是可导出的函数
- `injectors`
    - 注入器
    - 按依赖顺序调用 `providers` 的函数，该函数声明的主体是对 `wire.Build` 的调用
    - 编写注入器签名，然后 `wire` 命令生成函数
    - 函数的参数和返回值的类型必须唯一，负责不知道吧对应的值传递给谁

# 特殊用法

### `injector` 函数参数和返回值错误

最终的构造函数有多个返回值，我们也需要返回对应类型的值，值不重要，生成代码的时候会自动处理

### `ProviderSet` 分组

可以包含一组 `providers`， 使用 `wire.NewSet` 可以将多个 `providers` 构造函数合并到一个 `ProviderSet`

### `struct` 定制 `provider`

一个结构体很简单，可能不会为其添加一个构造函数，使用 `wire.Struct`

`func Struct(structType interface{}, fieldNames ...string) StructProvider`
需要显示的指定需要被赋值的字段，也可以使用通配符 `*`

### 使用 `struct` 字段作为 `provider`

使用 `wire.FieldsOf`

### 绑定之作为 `provider`

将一个值作为参数传递给 `wire.Value`

### 绑定接口作为 `provider`

`wire.InterfaceValue` 用法和 `wire.Value` 类似

### 绑定结构体到接口

`wire` 构建依靠参数，不支持接口类型
需要使用 `wire.Build` 告诉 `wire` 工具，将一个结构体绑定到接口

### 清理函数

函数的返回值可能包含一个清理函数，用来释放资源
清理函数的签名只能是 `func()`

### 注入器
返回值不重要，可以直接使用 `panic` 这样就不用写返回值
