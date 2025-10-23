package main

import (
	"io"
	"log"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func base() {
	ws := new(restful.WebService)

	ws.Route(ws.GET("/ping").To(pingPong))

	restful.Add(ws)

	http.ListenAndServe(":8000", nil)
}

func pingPong(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "pong")
}

func (u *UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUser).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").
			DataType("integer").
			DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(User{}).
		Returns(200, "OK", User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(u.createUser).
		Doc("create user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{}))

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		// docs
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

func userWeb() {
	uList := UserResource{map[string]User{}}
	restful.DefaultContainer.Add(uList.WebService())

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	log.Println("Start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	//base()
	userWeb()
}
