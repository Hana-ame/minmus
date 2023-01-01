package activityPub

import (
	"fmt"

	"github.com/gorilla/mux"
)

const debug = true

var Domain = "test.meromeromeiro.top"

func Main() {
	fmt.Println("23232423")
}

func RegisterHandlerFunc(r *mux.Router) {

	// Users
	r.HandleFunc("/users/{username}", Users)
	r.HandleFunc("/users/{username}/", Users)

	// Inbox
	r.HandleFunc("/users/{username}/inbox", Inbox)
	r.HandleFunc("/users/{username}/inbox/", Inbox)

	// SharedInbox
	r.HandleFunc("/inbox", Inbox)
	r.HandleFunc("/inbox/", Inbox)

}

// func
