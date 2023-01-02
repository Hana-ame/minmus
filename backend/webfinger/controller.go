package webfinger

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// 400 if the form mismatch
	if query["resource"] == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resources := query["resource"]

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
	if ok, _ := isExist(username); !ok {
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

	data, err := json.Marshal(resource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
