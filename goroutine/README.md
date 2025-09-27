# 获取当前 goroutine 的 id

没有直接暴露 `goroutine id` 的办法，有一些间接方法可以获取

### runtime.Stack 获取

获取当前 `goroutine` 的堆栈信息，通过解析堆栈信息可以获取
性能比较低

### 第三方库 goroutine id

`go get github.com/petermattis/goid`

使用了 `C` 和汇编来获取 `id`


