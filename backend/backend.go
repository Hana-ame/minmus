package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"

	"github.com/minmus/backend/general"
	"github.com/minmus/backend/webfinger"
)

func main() {
	general.Client = &http.Client{}

	color.Blue("starting")
	// activityPub.Main()

	r := mux.NewRouter()
	// r.NotFoundHandler = http.HandlerFunc(MyCustom404Handler)
	r.NotFoundHandler = http.HandlerFunc(wrapper(general.Proxy))

	r.HandleFunc(webfinger.WebFingerPath, wrapper(webfinger.Controller))
	r.HandleFunc("/test/", wrapper(general.Test))
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
