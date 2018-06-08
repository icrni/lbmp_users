package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
}

// Init router and start server
func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", greet).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", a.Router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
