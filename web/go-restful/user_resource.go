package main

import (
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"time"
)

type UserResource struct {
	users map[string]User
}

func (u *UserResource) findAllUser(req *restful.Request, resp *restful.Response) {
	log.Println("findAllUser")
	list := []User{}

	for _, each := range u.users {
		list = append(list, each)
	}

	resp.WriteEntity(list)
}

func (u *UserResource) findUser(req *restful.Request, resp *restful.Response) {
	log.Println("findUser")
	id := req.PathParameter("user-id")
	usr := u.users[id]
	if len(usr.ID) == 0 {
		resp.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		resp.WriteEntity(usr)
	}
}

func (u *UserResource) upsertUser(req *restful.Request, resp *restful.Response) {
	log.Println("upsertUser")
	usr := User{ID: req.PathParameter("user-id")}
	err := req.ReadEntity(&usr)
	if err == nil {
		u.users[usr.ID] = usr
		resp.WriteEntity(usr)
	} else {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	log.Println("createUser")
	usr := User{ID: fmt.Sprintf("%d", time.Now().Unix())}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.ID] = usr
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	log.Println("removeUser")
	id := request.PathParameter("user-id")
	delete(u.users, id)
}
