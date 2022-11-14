package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/test-web/middelware"
)

type controller struct{}

func NewController() *controller {
	return &controller{}
}

func (c *controller) RunController() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/token", GetToken).Methods("GET")
	router.HandleFunc("/login", middelware.IsAuthorized(Login)).Methods("POST")
	router.HandleFunc("/signin", GetToken).Methods("GET")
	router.HandleFunc("/students", middelware.IsAuthorized(GetAllStudents)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
