package main

import (
	"fmt"

	"github.com/pachirode/go-example/reflect/encoding/json"
)

type T1 struct {
	A int
	B string
	C int
}

func main() {
	t := T1{
		A: 1,
		B: "te",
		C: 2,
	}

	res, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Error marshal to JSON:", err)
		return
	}

	var d T1
	_ = json.Unmarshal(res, &d)
	fmt.Printf("%v", d)
}
