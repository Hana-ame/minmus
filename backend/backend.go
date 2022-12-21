package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"

	"github.com/hana-ame/minmus/backend/utils"
	"github.com/hana-ame/minmus/backend/webfinger"
)

var Domain string = "test.meromeromeiro.top"

func main() {
	utils.Client = &http.Client{}

	color.Blue("starting")
	// activityPub.Main()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(wrapper(utils.Proxy))

	r.HandleFunc(webfinger.WebFingerPath, wrapper(webfinger.Controller))

	// activityPub.RegisterHandlerFunc(r)

	r.HandleFunc("/test/", wrapper(utils.Test))
	r.HandleFunc("/users/{username}/inbox", wrapper(utils.Test))

	http.ListenAndServe("127.0.0.1:3001", r)
}

func MyCustom404Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	w.Header().Add("key", "value")
	w.WriteHeader(404)
	w.Header().Add("keykeykeykey", "valuevaluevaluevalue")
	w.Write([]byte(r.URL.String()))
}

func wrapper(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// something else
		// fmt.Println(r.URL.String())
		color.Cyan(r.URL.String())

		// origin funciton
		handler(w, r)
	}
}
