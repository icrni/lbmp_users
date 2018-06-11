package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/icrni/lbmp_users/app/user"
	"github.com/thedevsaddam/govalidator"

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
	a.Router.HandleFunc("/user", createUserHandlev2).Methods("POST")
}

func createUserHandle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t user.User
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	user.CreateUser(t.Firstname)
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func createUserHandlev2(w http.ResponseWriter, r *http.Request) {
	var t user.User
	rules := govalidator.MapData{
		"username": []string{"required", "between:5,15"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &t,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	t.CreateUser()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}
