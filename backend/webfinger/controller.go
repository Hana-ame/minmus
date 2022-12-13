package webfinger

import (
	"fmt"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	debug := true

	fmt.Println(r.URL.Path)
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
