package server

import (
	"net/http"

	"github.com/RushikeshMarkad16/Library-Managemant/api"
	"github.com/RushikeshMarkad16/Library-Managemant/user"
	"github.com/gorilla/mux"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User
	router.HandleFunc("/users", user.Create(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.FindByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.DeleteByID(dep.UserService)).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Update(dep.UserService)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
