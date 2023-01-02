package activityPub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hana-ame/minmus/backend/db"
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

}

func get(username string) (map[string]any, error) {

	arr := strings.Split(username, "@")
	r, err := http.Get(fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s", arr[len(arr)-1], username))
	if err != nil {
		return nil, err
	}
	var res webfinger.Resource
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	// links, ok := res["links"].([]interface{})
	// if !ok {
	// 	return nil, errors.New("not webfinger, find no links")
	// }

	// get id, id is the location of user object
	id, err := getIdFromWebfinger(res)
	if err != nil {
		return nil, err
	}
	// var id string
	// for _, link := range links {
	// 	if obj, ok := link.(map[string]interface{}); ok {
	// 		if rel, ok := obj["rel"]; ok {
	// 			if rel == "self" {
	// 				id = obj["href"]
	// 				break
	// 			}
	// 		}
	// 	}

	// }

}

func getIdFromWebfinger(res webfinger.Resource) (string, error) {
	for _, link := range res.Links {
		if link.Rel == "self" {
			return link.HRef, nil
		}
	}
	return "", fmt.Errorf("not webfinger, find no links with rel=self")
}
