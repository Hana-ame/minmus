package activityPub

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hana-ame/backend/general"
	"github.com/writeas/go-webfinger"
)

type resolver struct{}

func (resolve *resolver) FindUser(username string, hostname, requestHost string, r []webfinger.Rel) (*webfinger.Resource, error) {
	general.Test(nil, nil)
	fmt.Println(username, hostname, requestHost, r)
	return nil, nil
}

func (resolve *resolver) DummyUser(username string, hostname string, r []webfinger.Rel) (*webfinger.Resource, error) {
	fmt.Println(username, hostname, r)
	return nil, nil
}

func (resolve *resolver) IsNotFoundError(err error) bool {
	fmt.Println(err)
	return false
}

func Webfinger() func(w http.ResponseWriter, r *http.Request) {
	dummyResolver := &resolver{}
	// myResolver := webfinger.Resolver{dummyResolver}
	wf := webfinger.Default(dummyResolver)
	// wf.NotFoundHandler = // the rest of your app
	// wf.Webfinger(w, r)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		wf.Webfinger(w, r)
	}
}

// type webfinger struct {
// 	Subject string        `json:"subject"`
// 	Links   []interface{} `json:"links"`
// }

// func Webfinger(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	if debug {
// 		fmt.Println(query)
// 	}
// 	if query["resource"] == nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	resource := query["resource"]
// 	if debug {
// 		fmt.Println(resource)
// 	}

// 	// return

// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// }
