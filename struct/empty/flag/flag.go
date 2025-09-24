package main

// noCopy 的标准写法
type noCopy struct{}

// Lock/Unlock 方法是为了让 go vet 的 -copylocks 分析器识别
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type A struct {
	_ noCopy // 匿名字段即可，不需要名字
	a string
}

type B struct {
	b string
}

func main() {
	a := A{a: "hello"}
	b := B{b: "world"}

	// 普通使用不会触发警告
	_ = a
	_ = b

	// 触发 copy 检测
	c := a // 这里会拷贝 A，包含 noCopy 字段
	_ = c
}
