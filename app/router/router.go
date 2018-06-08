package router

import (
	"github.com/gorilla/mux"
)

Router = mux.NewRouter()

func Init(router *router) {
	router = mux.NewRouter()
}

func setRoutes(r *mux.Router) {
	r.HandleFunc("/", hello).Methods("GET")
}

func hello() {

}
