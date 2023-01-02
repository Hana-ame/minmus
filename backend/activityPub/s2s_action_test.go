package activityPub

import (
	"fmt"
	"testing"

	"github.com/hana-ame/minmus/backend/db"
)

func Test1(t *testing.T) {
	m := map[string]interface{}{
		"a": nil,
	}
	fmt.Println(m)
	fmt.Println(m["a"])

	s, ok := m["a"].(string)
	fmt.Println(s, ok)
}

func TestGetUserActivityStreamFromRemote(t *testing.T) {
	m, err := getUserActivityStreamFromRemote("meiro@misskey.meromeromeiro.top")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}

func TestGetUserFromRemote(t *testing.T) {
	u, err := getUserFromRemote("meiro@misskey.meromeromeiro.top")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}

func TestGetRemoteUser(t *testing.T) {
	db.InitDB()
	u, err := getRemoteUser("meiro@misskey.meromeromeiro.top")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}
