package main

import (
	"flag"
	"fmt"
)

type flagVal struct {
	val string
}

func (v *flagVal) String() string {
	return v.val
}

func (v *flagVal) Set(s string) error {
	v.val = s
	return nil
}

func base() {

	// flag.Type() 返回
	var intFlag = flag.Int("int", 12, "int msg")

	// flag.TypeVar() 绑定命令行参数
	var flagVar int
	flag.IntVar(&flagVar, "flagvar", 123, "int var msg")

	// flag.Var() 绑定命令行参数
	val := flagVal{}
	flag.Var(&val, "val", "val msg")

	flag.Parse()

	fmt.Println(intFlag)
	fmt.Println(flagVar)
	fmt.Printf("%v\n", val)
	fmt.Println("===========================")

	fmt.Printf("NFlag: %v\n", flag.NFlag()) // 返回已设置的命令行标志个数
	fmt.Printf("NArg: %v\n", flag.NArg())   // 返回处理完标志后剩余的参数个数
	fmt.Printf("Args: %v\n", flag.Args())   // 返回处理完标志后剩余的参数列表
	fmt.Printf("Arg(1): %v\n", flag.Arg(1)) // 返回处理完标志后剩余的参数列表中第 i 项
}

func main() {
	base()
}
