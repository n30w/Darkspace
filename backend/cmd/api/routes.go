package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	// Enhanced routing patterns in Go 1.22. HandleFuncs now
	// accept a method and a route variable parameter.
	// https://tip.golang.org/doc/go1.22

	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc("GET /v1/home", app.homeHandler)
	router.HandleFunc("GET /v1/course/{id}", app.courseHomepageHandler)

	// Course CRUD operations
	router.HandleFunc("/v1/course/create", app.courseCreateHandler)
	router.HandleFunc("/v1/course/read", app.courseReadHandler)
	router.HandleFunc("/v1/course/update", app.courseUpdateHandler)
	router.HandleFunc("/v1/course/delete", app.courseDeleteHandler)

	// User CRUD operations
	router.HandleFunc("/v1/user/create", app.userCreateHandler)
	router.HandleFunc("/v1/user/read", app.userReadHandler)
	router.HandleFunc("/v1/user/update", app.userUpdateHandler)
	router.HandleFunc("/v1/user/delete", app.userDeleteHandler)

	// A user posts something to a discussion
	router.HandleFunc("/v1/user/post", app.userPostHandler)

	// Login will require authorization, body will contain the credential info
	router.HandleFunc("/v1/user/login", app.userLoginHandler)

	// Assignment CRUD operations
	router.HandleFunc("/v1/course/assignment/create", app.assignmentCreateHandler)
	router.HandleFunc("/v1/course/assignment/read", app.assignmentReadHandler)
	router.HandleFunc("/v1/course/assignment/update", app.assignmentUpdateHandler)
	router.HandleFunc("/v1/course/assignment/delete", app.assignmentDeleteHandler)

	// Discussion CRUD operations
	router.HandleFunc("/v1/course/discussion/create", app.discussionCreateHandler)
	router.HandleFunc("/v1/course/discussion/read", app.discussionReadHandler)
	router.HandleFunc("/v1/course/discussion/update", app.discussionUpdateHandler)
	router.HandleFunc("/v1/course/discussion/delete", app.discussionDeleteHandler)

	return router
}
