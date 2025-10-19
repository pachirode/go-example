package main

import (
	"errors"
	"fmt"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func base() {
	var errs []error

	errs = append(errs, errors.New("1"))
	errs = append(errs, errors.New("2"))

	agg := utilerrors.NewAggregate(errs)

	fmt.Printf("errs len: %d, aggregate len: %d\n", len(errs), len(agg.Errors()))

	fmt.Printf("errs: %s\n", errs)
	fmt.Printf("errors: %s\n", agg.Errors())

	fmt.Printf("aggregate: %s\n", agg)
	fmt.Printf("err1: %s, err2: %s\n", agg.Errors()[0], agg.Errors()[1])
}

func main() {
	base()
}
