package webfinger

import (
	"fmt"
	"strings"

	"github.com/hana-ame/minmus/backend/utils"
)

// this username shall not contain @[domain]
func isExist(username string) bool {
	// TODO: should connect to database to check if username exists

	return true
}

func getResource(username string) *Resource {
	// return mockGet(username)

	// TODO: sould connect to database to find Aliases
	self := &Link{
		Rel:  "self",
		Type: "application/activity+json",
		HRef: fmt.Sprintf("https://%s/users/%s", Domain, username),
	}
	// porfile
	profilePage := &Link{
		Rel:  "http://webfinger.net/rel/profile-page",
		Type: "text/html",
		HRef: fmt.Sprintf("https://%s/@%s", Domain, username),
	}
	// what's this used for?
	subscribe := &Link{
		Rel:      "http://ostatus.org/schema/1.0/subscribe",
		Template: fmt.Sprintf("https://%s/authorize-follow?acct={uri}", Domain),
	}

	r := &Resource{
		Subject: fmt.Sprintf("acct:%s@%s", username, Domain),
		Aliases: nil, // TODO
		Links: []Link{
			*self,
			*profilePage,
			*subscribe,
		},
	}
	return r
	// return mockGet(username)
}

// for test
func mockGet(username string) *Resource {
	data := utils.Get("https://misskey.meromeromeiro.top/.well-known/webfinger?resource=acct:" + username + "@misskey.meromeromeiro.top")
	if data == nil {
		return nil
	}

	r := &Resource{}

	r = utils.Unmarshal(data, r)
	if r == nil {
		return r
	}

	r.Subject = strings.ReplaceAll(r.Subject, "misskey.meromeromeiro.top", Domain)
	for k := range r.Links {
		r.Links[k].HRef = strings.ReplaceAll(r.Links[k].HRef, "misskey.meromeromeiro.top", Domain)
	}

	return r
}
