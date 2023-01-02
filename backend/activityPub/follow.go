package activityPub

import (
	"fmt"
	"strconv"

	"github.com/hana-ame/minmus/backend/db"
)

var FollowItemsPerPage = 10

func getActivityStreamOfFollowers(username string, page int) (map[string]any, error) {
	id := fmt.Sprintf("https://%s/users/%s/followers", Domain, username)
	as := make(map[string]any)
	as["@context"] = "https://www.w3.org/ns/activitystreams"
	as["type"] = "OrderedCollectionPage"
	as["partOf"] = id
	as["id"] = id + "?page=" + strconv.Itoa(page)

	// do query
	user := &db.User{Username: username}
	user, err := db.QueryUser(user)
	if err != nil {
		return nil, err
	}
	followersCount := user.FollowersCount

	as["totalItems"] = followersCount

	// outOfIdx
	if page > (followersCount-1)/FollowItemsPerPage+1 {
		as["orderedItems"] = []any{}
		return as, nil
	}
	if page > 1 {
		as["prev"] = id + "?page=" + strconv.Itoa(page-1)
	}
	if page < (followersCount-1)/FollowItemsPerPage+1 {
		as["next"] = id + "?page=" + strconv.Itoa(page+1)
	}

	// do query TODO
	as["orderedItems"] = []any{}

	return as, nil
}
