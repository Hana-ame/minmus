package activityPub

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
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
	// fmt.Println("path", username, ok)

	// fmt.Println(r.Header["Accept"])

	// if accept not match return 302
	// if accept := r.Header.Get("Accept"); accept != "application/activity+json" && accept != "application/ld+json" {
	if accept := r.Header.Get("Accept"); strings.HasPrefix(accept, "application/activity+json") && strings.HasPrefix(accept, "application/ld+json") {
		http.Redirect(w, r, fmt.Sprintf("https://%s/@%s", Domain, username), http.StatusFound)
	}

	// get user (type Person)
	person := GetPerson(username)
	if person == nil {
		return
	}

	data := utils.Marshal(person)

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// TODO: connect to db for data
func GetPerson(username string) activityStream {
	return dummyPerson(username)
}

func dummyPerson(username string) activityStream {
	// user := &db.User{
	// 	Username: "dummy",
	// }
	// _, err := db.QueryUser(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	as := make(activityStream, 0)
	as["@context"] = []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}
	as["type"] = "Person"
	as["id"] = getID(Domain, username)
	as["inbox"] = getInbox(Domain, username)
	as["outbox"] = getOutbox(Domain, username)
	as["followers"] = getFollowers(Domain, username)
	as["following"] = getFollowing(Domain, username)
	as["featured"] = getFeatured(Domain, username)
	as["sharedInbox"] = getSharedInbox(Domain)
	as["endpoints"] = activityStream{
		"sharedInbox": getSharedInbox(Domain),
	}

	as["url"] = getURL(Domain, username)
	// dummy
	as["preferredUsername"] = "username"
	as["name"] = nil
	as["summary"] = nil
	as["icon"] = nil
	as["image"] = nil
	as["tag"] = []string{}
	as["manuallyApprovesFollowers"] = false
	as["publicKey"] = getPublicKey(Domain, username)

	return as
}

func getID(domain string, username string) string {
	return fmt.Sprintf("https://%s/users/%s", domain, username)
}
func getInbox(domain string, username string) string {
	return getID(domain, username) + "/inbox"
}
func getOutbox(domain string, username string) string {
	return getID(domain, username) + "/outbox"
}
func getFollowers(domain string, username string) string {
	return getID(domain, username) + "/followers"
}
func getFollowing(domain string, username string) string {
	return getID(domain, username) + "/following"
}
func getFeatured(domain string, username string) string {
	return getID(domain, username) + "/collections/featured"
}
func getSharedInbox(domain string) string {
	return fmt.Sprintf("https://%s/inbox", Domain)
}
func getURL(domain string, username string) string {
	return fmt.Sprintf("https://%s/@%s", Domain, username)
}
func getPublicKey(domain string, username string) activityStream {

	user := &db.User{
		Username: "dummy",
	}
	_, err := db.QueryUser(user)
	if err != nil {
		log.Fatal(err)
	}

	as := make(activityStream, 0)

	as["id"] = getID(domain, username) + "#main-key"
	as["type"] = "Key"
	as["owner"] = getID(domain, username)
	as["publicKeyPem"] = user.PublicKeyPem

	return as
}
