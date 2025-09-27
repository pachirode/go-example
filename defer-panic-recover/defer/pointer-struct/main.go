package main

import "fmt"

type User struct {
	name string
}

func (u User) Name() {
	fmt.Println("Name:", u.name)
}

func (u *User) PointName() {
	fmt.Println("PointName:", u.name)
}

func printUser() {
	u := User{name: "user1"}

	defer u.Name()
	defer u.PointName()

	u.name = "user2"
}

func printUser2() {
	u := User{name: "user1"}

	defer func() {
		u.Name()
		u.PointName()
	}()

	u.name = "user2"
}

func main() {
	printUser()
	printUser2()
}
