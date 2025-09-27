package main

import "fmt"

func g(i int) (number int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("f returns err: %v", r)
			}
		}
	}()

	number, err = f(i)
	return number, err
}

func f(i int) (int, error) {
	if i == 0 {
		panic("i=0")
	}
	return i * i, nil
}

func main() {
	fmt.Println(g(1))
	fmt.Println(g(0))
}
