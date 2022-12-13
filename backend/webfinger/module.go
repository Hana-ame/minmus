package webfinger

import (
	"github.com/minmus/backend/general"
)

func mockGet(username string) *Resource {
	data := general.Get("https://misskey.meromeromeiro.top/.well-known/webfinger?resource=acct:" + username + "@misskey.meromeromeiro.top")
	if data == nil {
		return nil
	}

	r := &Resource{}

	r = general.Unmarshal(data, r)

	return r
}
