package activityPub

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// /users/{username}/
// /users/{username}
func Users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	fmt.Println("path", username, ok)

	fmt.Println(r.Header["Accept"])

	// if accept not match return 302
	if accept := r.Header.Get("Accept"); accept != "application/activity+json" && accept != "application/ld+json" {
		http.Redirect(w, r, fmt.Sprintf("https://%s/@%s", Domain, username), http.StatusFound)
	}

	// get user (type Person)
	person := GetPerson(username)
	if person == nil {
		return
	}
}

// TODO: connect to db for data
func GetPerson(username string) activityStream {
	return dummyPerson()
}

func dummyPerson() activityStream {
	return nil
}
