package main

type User struct {
	ID   string `json:"ID" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user"`
}
