# `pkg/errors`

### fundamental

- `msg`
    - 记录错误信息
- `*stack`
    - 指针字段，用来记录出错时的错误堆栈信息
- 输出格式
    - `Format` 参数实现了 `io.Writer` 接口，该参数写的内容都会被格式化
    - `%s`
    - `%v`
    - `%+v`
        - 错误栈和堆栈信息
    - `%q`

### withStack

- `error`
    - 错误
- `*stack`
    - 错误堆栈
- `Cause`
    - 表示错误根因

### withMessage

给一个错误附加一个新的错误信息

# stack

### callers

底层调用了 `runtime.Callers` 获取当前调用栈信息

# `func Callers(skip int, pc []uintptr) int`

- `skip`
    - 指定要跳过的栈帧的数量
    - 0
        - 包含 `runtime.Callers` 调用帧
    - 1
        - 从调用 `runtime.Callers` 函数开始
- `pc []uintptr`
    - 会填充传递过来的指针切片
    - `goroutine` 调用栈的程序器（PC）通过这个值可以获取更多详细信息
- 返回值
    - 记录填充到 `pc` 中的切片数量，如果实际调用小于 `pc`，就返回小的

# 错误处理

### Sentinel error

预定义的错误被称为哨兵错误，大部分被定义为包级别公开变量，以 `Err` 开头

### 常量错误

无法将 `errors.New` 的返回值赋值给一个变量，需要使用自定义 `error` 类型来解决这个问题

### 错误值比较

出现多种错误，可以使用 `switch...case...` 来判定错误值

### 类型断言

需要导致指定的错误类型，使得代码存在较强的依赖性

- `Type Assertion`
    - `error` 也是一个普通的接口
- `Type Switch`
    - 禁止使用 `fallthrough` 关键字

### 行为断言

断言错误行为

```text
// net.Error

type Error interface {
	error
	Timeout() bool // Is the error a timeout?

	// Deprecated: Temporary errors are not well-defined.
	// Most "temporary" errors are timeouts, and the few exceptions are surprising.
	// Do not use this method.
	Temporary() bool
}
```

### 暂存错误状态

链式调用或者循环等场景下，暂存中间过程出现的全部错误，只用在最后一次处理错误，调用链中出现错误，后面都可以自行处理错误

### 返回错误而不是指针

### 不能忽略掉错误

如果确幸不会返回错误，需要在函数内部处理完 `MustXXX` 在 `Gin` 比较常见

### nil 错误值不等于 nil

一个接口对象实际包含两个属性

- 类型
    - `T`
- 具体值
    - `V`

只有当 `T=nil, V 没有设置` 接口的值才为 `nil`