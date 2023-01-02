package activityPub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

// /users/{username}/
// /users/{username}
func Users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// check the header, if not s2s request, redirect to normal profile
	if accept := r.Header.Get("Accept"); !(strings.HasPrefix(accept, "application/activity+json") || strings.HasPrefix(accept, "application/ld+json")) {
		http.Redirect(w, r, fmt.Sprintf("/@%s", username), http.StatusFound)
	}

	// get user (type *db.User)
	user, err := getLocalUserByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// it will no go error?
	as, err := convertLocalUserToS2SActivityStream(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// data := utils.Marshal(person)
	data, err := json.Marshal(as)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/activity+json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// /{username}/inbox
// /{username}/inbox/
func Inbox(w http.ResponseWriter, r *http.Request) {
	// only POST
	var err error
	color.Green(fmt.Sprint(r)) // debug

	r.Host = Domain

	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// redirect to profile page if not s2s action
	if r.Method != "POST" {
		http.Redirect(w, r, fmt.Sprintf("/@%s", username), http.StatusMovedPermanently)
	}

	// httpsig
	err = verify(r)
	if err != nil {
		color.Red(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// verify success

	text, err := io.ReadAll(r.Body)
	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Yellow(string(text))

}

// /inbox
// /inbox/
func SharedInbox(w http.ResponseWriter, r *http.Request) {
	var err error
	color.Green(fmt.Sprint(r))
	text, err := io.ReadAll(r.Body)
	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Yellow(string(text))
	return
}
