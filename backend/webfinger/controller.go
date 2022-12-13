package webfinger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/minmus/backend/general"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	debug := true

	if debug {
		fmt.Println(r.URL.Path)
	}

	query := r.URL.Query()
	if debug {
		fmt.Println(query)
	}

	if query["resource"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resources := query["resource"]
	if debug {
		fmt.Println(resources)
	}

	if len(resources) != 1 || !strings.HasPrefix(resources[0], "acct:") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, domain := resolveAcct(resources[0])
	if domain != Domain {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	resource := mockGet(username)
	fmt.Println(resource)
	if resource == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resource.Subject = strings.ReplaceAll(resource.Subject, "misskey.meromeromeiro.top", Domain)
	for k := range resource.Links {
		resource.Links[k].HRef = strings.ReplaceAll(resource.Links[k].HRef, "misskey.meromeromeiro.top", Domain)
	}

	data := general.Marshal(*resource)

	// return

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(data)
}

func resolveAcct(resource string) (string, string) {
	arr := strings.Split(resource[5:], "@")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return resource, ""
}
