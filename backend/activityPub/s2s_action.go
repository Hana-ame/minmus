package activityPub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
	"github.com/hana-ame/minmus/backend/webfinger"
)

func Follow() {}

func getRemoteUser(username string) (*db.User, error) {
	// first check if it is in db
	obj := &db.User{
		Username: username,
	}
	user, err := db.QueryUser(obj)
	if err == nil {
		return user, nil
	}

	// not found, query from other site
	user, err = getUserFromRemote(username)
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

func getUserFromRemote(username string) (*db.User, error) {

	obj, err := getUserActivityStreamFromRemote(username)
	if err != nil {
		return nil, err
	}

	// from map[string]any to struct
	user, err := getUserFromActivityStream(obj)
	if err != nil {
		return nil, err
	}
	user.Username = username

	return user, nil
}

func getUserActivityStreamFromRemote(username string) (map[string]any, error) {

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

func getUserFromActivityStream(m map[string]any) (*db.User, error) {
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
