package webfinger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hana-ame/minmus/backend/utils"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	debug := false

	if debug {
		fmt.Println(r.URL.Path)
	}

	query := r.URL.Query()
	if debug {
		fmt.Println(query)
	}

	// 400 if the form mismatch
	if query["resource"] == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resources := query["resource"]
	if debug {
		fmt.Println(resources)
	}

	// 400 if the form mismatch
	if len(resources) != 1 || !strings.HasPrefix(resources[0], "acct:") {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// 422 if domain not match
	username, domain := resolveAcct(resources[0])
	if domain != Domain {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// 404 if username not exist
	if !isExist(username) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	resource := getResource(username)
	// fmt.Println(resource)
	// 404 if username not exist
	if resource == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	data := utils.Marshal(*resource)

	// return
	w.Header().Set("content-type", "application/jrd+json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func resolveAcct(resource string) (string, string) {
	arr := strings.Split(resource[5:], "@")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return resource, ""
}
