package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/icrni/lbmp_users/app/user"

	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
}

// Init router and start server
func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.setRoutes()
	log.Fatal(http.ListenAndServe(":3000", a.Router))
}

func (a *App) setRoutes() {
	a.Router.HandleFunc("/", greet).Methods("GET")
}

func greet(w http.ResponseWriter, r *http.Request) {
	user.CreateUser()
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
