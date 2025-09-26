package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func a() error {
	return errors.New("a error")
}

func b() error {
	err := a()
	if err != nil {
		newErr := errors.New("b error")
		//newErr := errors.Wrap(err, "b error")
		return newErr
	}

	return nil
}

func main() {
	err := b()
	if err != nil {
		// %v 打印错误信息
		fmt.Printf("%v\n", err)

		fmt.Println("============================================")

		// %+v 打印错误信息和错误堆栈
		fmt.Printf("%+v\n", err)

		fmt.Println("============================================")

		// 打印错误根因
		fmt.Printf("%v\n", errors.Cause(err))
		return
	}
	fmt.Println("success")
}
