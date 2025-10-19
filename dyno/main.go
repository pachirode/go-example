package main

import (
	"fmt"
	"github.com/icza/dyno"
)

func get() {
	y := map[string]interface{}{
		"object": map[interface{}]interface{}{
			"a": 1,
			"array": []interface{}{
				map[string]interface{}{
					"null_value": interface{}(nil),
				},
				map[string]interface{}{
					"boolean": true,
				},
				map[string]interface{}{
					"integer": 1,
				},
			},
			"key": "value",
			1:     2,
		},
	}

	get, err := dyno.Get(y, "object", "array", 2, "integer")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T, %#v\n", get, get)
}

func set() {
	m := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "test",
			"address": map[string]interface{}{
				"city": "ShangHai",
			},
		},
	}

	err := dyno.Set(m, "city", "user", "address", "city")
	if err != nil {
		panic(err)
	}

	val, _ := dyno.Get(m, "user", "address", "city")

	fmt.Printf("%v", val)
}

func main() {
	//get()
	set()
}
