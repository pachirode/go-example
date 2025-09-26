package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	{
		name := "a"

		val := reflect.ValueOf(name)
		typ := reflect.TypeOf(name)
		fmt.Println(val, typ)

		{
			user := User{
				Name: "t",
			}

			val := reflect.ValueOf(user)
			typ := reflect.TypeOf(user)
			fmt.Println(val, typ)
		}
	}

	{
		user := &User{
			Name: "d",
		}
		kind := reflect.ValueOf(user).Kind()
		fmt.Println(kind)
		kind = reflect.ValueOf(*user).Kind()
		fmt.Println(kind)
		kind = reflect.ValueOf(user).Elem().Kind()
		fmt.Println(kind)
	}

	{
		user := &User{
			Name: "d",
		}
		nameField := reflect.ValueOf(user).Elem().FieldByName("Name")
		nameField = reflect.ValueOf(user).Elem().FieldByIndex([]int{0})
		nameField = reflect.ValueOf(user).Elem().Field(0)
		fmt.Println(nameField)
	}

	{
		user := &User{
			Name: "d",
		}
		nameField, _ := reflect.TypeOf(user).Elem().FieldByName("Name")
		fmt.Printf("%+v\n", nameField)
	}

}
