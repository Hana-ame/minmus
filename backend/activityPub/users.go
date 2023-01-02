package activityPub

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
	"github.com/hana-ame/minmus/backend/webfinger"
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

func getRemoteUserByName(username string) (*db.User, error) {
	// first check if it is in db
	obj := &db.User{
		Username: username,
	}
	user, err := db.QueryUser(obj)
	if err == nil {
		return user, nil
	}

	// not found, query from other site
	user, err = fetchUserFromRemote(username)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	// put into db
	err = db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	// done?
	return user, nil
}

func fetchUserFromRemote(username string) (*db.User, error) {

	obj, err := fetchUserActivityStreamFromRemote(username)
	if err != nil {
		return nil, err
	}

	// from map[string]any to struct
	user, err := convertActivityStreamToRemoteUser(obj)
	if err != nil {
		return nil, err
	}
	user.Username = username

	return user, nil
}

func fetchUserActivityStreamFromRemote(username string) (map[string]any, error) {

	arr := strings.Split(username, "@")
	r, err := http.Get(fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s", arr[len(arr)-1], username))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var res webfinger.Resource
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	// get id, id is the location of user object
	id, err := webfinger.GetIdFromWebfinger(res)
	if err != nil {
		return nil, err
	}

	// get user object from other site
	r, err = utils.GetWithAccept(id)
	if err != nil {
		return nil, err
	}

	var obj map[string]any
	err = json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func convertActivityStreamToRemoteUser(m map[string]any) (*db.User, error) {
	// id
	id, ok := m["id"].(string)
	if !ok {
		return nil, fmt.Errorf("user id is not string")
	}
	// endpoints
	inbox, ok := m["inbox"].(string)
	if !ok {
		return nil, fmt.Errorf("user inbox is not string")
	}
	outbox, ok := m["outbox"].(string)
	if !ok {
		return nil, fmt.Errorf("user outbox is not string")
	}
	followers, ok := m["followers"].(string)
	if !ok {
		return nil, fmt.Errorf("user followers is not string")
	}
	following, ok := m["following"].(string)
	if !ok {
		return nil, fmt.Errorf("user following is not string")
	}
	featured, ok := m["featured"].(string)
	if !ok {
		return nil, fmt.Errorf("user featured is not string")
	}
	sharedInbox, ok := m["sharedInbox"].(string)
	if !ok {
		return nil, fmt.Errorf("user sharedInbox is not string")
	}
	// public key
	publicKey, ok := m["publicKey"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("user publicKey is not map[string]any")
	}
	publicKeyID, ok := publicKey["id"].(string)
	if !ok {
		return nil, fmt.Errorf("publicKey id is not string")
	}
	publicKeyPem, ok := publicKey["publicKeyPem"].(string)
	if !ok {
		return nil, fmt.Errorf("publicKey public key pem is not string")
	}
	// user info
	preferredUsername, ok := m["preferredUsername"].(string)
	if !ok {
		return nil, fmt.Errorf("user preferredUsername is not string")
	}
	name, ok := m["name"].(string)
	if !ok {
		// return nil, fmt.Errorf("user name is not string")
	}
	summary, ok := m["summary"].(string)
	if !ok {
		// return nil, fmt.Errorf("user summary is not string")
	}
	// url, ok := m["url"].(string)
	// if !ok{
	// 	return nil, fmt.Errorf("user url is not string")
	// }
	icon, ok := m["icon"].(string)
	if !ok {
		if obj, ok := m["icon"].(map[string]any); ok {
			if icon, ok = obj["url"].(string); !ok {
				// return nil, fmt.Errorf("icon url is not string")
			}
		} else {
			// return nil, fmt.Errorf("icon is not string or map[string]any")
		}
	}
	image, ok := m["image"].(string)
	if !ok {
		if obj, ok := m["image"].(map[string]any); ok {
			if image, ok = obj["url"].(string); !ok {
				// return nil, fmt.Errorf("icon url is not string")
			}
		} else {
			// return nil, fmt.Errorf("icon is not string or map[string]any")
		}
	}

	return &db.User{
		// where to get this obj
		ID: id,
		// endpoints
		Inbox:       inbox,
		Outbox:      outbox,
		Followers:   followers,
		Following:   following,
		Featured:    featured,
		SharedInbox: sharedInbox,
		// public key
		PublicKeyID:  publicKeyID,
		PublicKeyPem: publicKeyPem,
		// user info
		PreferredUsername: preferredUsername,
		Name:              name,
		Summary:           summary,
		Icon:              icon,
		Image:             image,
	}, nil
}

// no longer used
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
