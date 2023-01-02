package activityPub

import (
	"fmt"
	"log"

	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
)

// TODO: connect to db for data
func GetPerson(username string) map[string]any {
	return dummyPerson(username)
}

func getLocalUserByName(username string) (*db.User, error) {
	user := &db.User{
		Username: username,
	}
	_, err := db.QueryUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func convertLocalUserToS2SActivityStream(user *db.User) (map[string]any, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}
	username := user.Username
	if !utils.IsValidUsername(username) {
		return nil, fmt.Errorf("not valid username")
	}

	id := fmt.Sprintf("https://%s/users/%s", Domain, username)
	sharedInbox := fmt.Sprintf("https://%s/inbox", Domain)
	// pubKeyID =
	preferredUsername := username
	var name any = user.Name
	if name == "" {
		name = nil
	}
	var summary any = user.Summary
	if summary == "" {
		summary = nil
	}
	var icon any = user.Icon
	if icon == "" {
		icon = nil
	}
	var image any = user.Image
	if image == "" {
		image = nil
	}
	publicKeyPem := user.PublicKeyPem
	manuallyApprovesFollowers := user.ManuallyApprovesFollowers

	as := make(map[string]any, 0)
	as["@context"] = []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}
	as["type"] = "Person"
	as["id"] = id
	as["inbox"] = id + "/inbox"
	as["outbox"] = id + "/outbox"
	as["followers"] = id + "/followers"
	as["following"] = id + "/following"
	as["featured"] = id + "/collections/featured"
	as["sharedInbox"] = sharedInbox
	as["endpoints"] = map[string]any{
		"sharedInbox": sharedInbox,
	}

	as["url"] = getURL(Domain, username)
	// dummy
	as["preferredUsername"] = preferredUsername
	as["name"] = name
	as["summary"] = summary
	as["icon"] = icon
	as["image"] = image
	as["tag"] = []string{}
	as["manuallyApprovesFollowers"] = manuallyApprovesFollowers
	as["publicKey"] = map[string]any{
		"id":           id + "#main-key",
		"type":         "Key",
		"owner":        id,
		"publicKeyPem": publicKeyPem,
	}

	return as, nil
}

func dummyPerson(username string) map[string]any {
	// user := &db.User{
	// 	Username: "dummy",
	// }
	// _, err := db.QueryUser(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	as := make(map[string]any, 0)
	as["@context"] = []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}
	as["type"] = "Person"
	as["id"] = getID(Domain, username)
	as["inbox"] = getInbox(Domain, username)
	as["outbox"] = getOutbox(Domain, username)
	as["followers"] = getFollowers(Domain, username)
	as["following"] = getFollowing(Domain, username)
	as["featured"] = getFeatured(Domain, username)
	as["sharedInbox"] = getSharedInbox(Domain)
	as["endpoints"] = map[string]any{
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

	// fmt.Println(as)
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
func getPublicKey(domain string, username string) map[string]any {

	user := &db.User{
		Username: "dummy",
	}
	_, err := db.QueryUser(user)
	if err != nil {
		log.Fatal(err)
	}

	as := make(map[string]any, 0)

	as["id"] = getID(domain, username) + "#main-key"
	as["type"] = "Key"
	as["owner"] = getID(domain, username)
	as["publicKeyPem"] = user.PublicKeyPem

	return as
}
