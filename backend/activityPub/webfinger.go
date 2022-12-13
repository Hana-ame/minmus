package activityPub

import (
	"fmt"
	"net/http"
)

type webfinger struct {
	Subject string        `json:"subject"`
	Links   []interface{} `json:"links"`
}

func Webfinger(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if debug {
		fmt.Println(query)
	}
	if query["resource"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resource := query["resource"]
	if debug {
		fmt.Println(resource)
	}

	// return

	w.Header().Set("Access-Control-Allow-Origin", "*")

}
