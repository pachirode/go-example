package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func base() {
	var errs *multierror.Error

	errs = multierror.Append(errs, errors.New("err1"))
	errs = multierror.Append(errs, errors.New("err2"))

	fmt.Printf("%v", errs)

	// 修改输出格式
	errs.ErrorFormat = func(errs []error) string {
		var res string
		for i := 0; i < len(errs); i++ {
			res = res + errs[i].Error() + "~~~~~~~~~~~~"
		}
		return res
	}
	fmt.Println(errs.Error())
}

func main() {
	base()
}
