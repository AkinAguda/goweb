package main

import (
	"log"
	"net/http"

	login "./Handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var person login.User
	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.GET("/confirm/:token", person.ConfirmToken)
	router.GET("/signup", person.SignUp)
	router.POST("/signup/", person.SignUpPost)
	router.GET("/login", person.Login)
	router.POST("/login/", person.LoginPost)
	router.GET("/", person.Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
