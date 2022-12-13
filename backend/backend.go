package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/writeas/go-webfinger"

	"github.com/minmus/backend/activityPub"
)

func main() {
	activityPub.Main()

	r := mux.NewRouter()

	r.HandleFunc("/", scope)
	r.HandleFunc(webfinger.WebFingerPath, activityPub.Webfinger())
	http.ListenAndServe("127.0.0.1:3001", r)
}

func scope(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Add("key", "value")
	w.WriteHeader(200)
	w.Header().Add("keykeykeykey", "valuevaluevaluevalue")
	w.Write([]byte("hello\n"))
}
