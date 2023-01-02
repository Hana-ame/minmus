package activityPub

import (
	"fmt"
	"testing"

	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
)

func TestUser(t *testing.T) {
	var as map[string]any
	var data []byte
	for i := 0; i < 10_000; i++ {
		as = dummyPerson("username111")
		data = utils.Marshal(as)
	}
	fmt.Println(as)
	fmt.Println(string(data))
}

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
	m, err := fetchUserActivityStreamFromRemote("meiro@misskey.meromeromeiro.top")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}

func TestGetUserFromRemote(t *testing.T) {
	u, err := fetchUserFromRemote("meiro@misskey.meromeromeiro.top")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}

func TestGetRemoteUser(t *testing.T) {
	db.InitDB()
	u, err := getRemoteUserByName("meiro@misskey.meromeromeiro.top")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}
