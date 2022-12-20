package webfinger

import (
	"github.com/minmus/backend/general"
)

// this username shall not contain @[domain]
func isExist(username string) bool {
	// TODO: should connect to database
	return true
}

func getResource(username string) *Resource {
	return mockGet(username)
}

func mockGet(username string) *Resource {
	data := general.Get("https://misskey.meromeromeiro.top/.well-known/webfinger?resource=acct:" + username + "@misskey.meromeromeiro.top")
	if data == nil {
		return nil
	}

	r := &Resource{}

	r = general.Unmarshal(data, r)

	return r
}
