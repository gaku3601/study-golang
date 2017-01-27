package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", GetHelloWorld),
		rest.Get("/#param", GetHelloWorldParam),
		rest.Post("/", PostHelloWorld),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type User struct {
	Name     string
	Password string
}

func GetHelloWorld(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Body": "Hello World!"})
}

func GetHelloWorldParam(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Hello": r.PathParam("param")})
}

func PostHelloWorld(w rest.ResponseWriter, r *rest.Request) {
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Name == "" {
		rest.Error(w, "user name required", 400)
		return
	}
	if user.Password == "" {
		rest.Error(w, "user password required", 400)
		return
	}
	w.WriteJson(&user)
	w.WriteJson(map[string]string{"gakusName": user.Name})
}
