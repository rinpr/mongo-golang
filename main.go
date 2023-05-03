package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"mongo-golang/controllers"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Println("cant listen and serve")
		fmt.Println(err)
	}
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("cant connect to db")
		fmt.Println(err)
	}
	return session
}
