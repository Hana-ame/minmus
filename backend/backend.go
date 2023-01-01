package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"

	"github.com/hana-ame/minmus/backend/activityPub"
	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
	"github.com/hana-ame/minmus/backend/webfinger"
)

var Domain string = "meiro.meromeromeiro.top"

func main() {

	activityPub.Domain = Domain
	webfinger.Domain = Domain

	db.InitDB()

	utils.Client = &http.Client{}

	color.Blue("starting")
	// activityPub.Main()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(wrapper(MyCustom404Handler))

	r.HandleFunc(webfinger.WebFingerPath, wrapper(webfinger.Controller))

	activityPub.RegisterHandlerFunc(r)

	r.HandleFunc("/test/", wrapper(utils.Test))
	// r.HandleFunc("/users/{username}/inbox", wrapper(utils.Test))

	http.ListenAndServe("127.0.0.1:3001", r)
}

func MyCustom404Handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	color.Green(fmt.Sprint(r))
	text, err := io.ReadAll(r.Body)
	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Yellow(string(text))
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
