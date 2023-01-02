package webcore

import (
	"testing"

	"github.com/hana-ame/minmus/backend/db"
)

func TestCreateUser(t *testing.T) {
	db.InitDB()
	var err error

	err = CreateUser("user123", "pass123", "")
	if err != nil {
		t.Error(err)
	}
	err = CreateUser("useR123", "pass123", "")
	if err != nil {
		// t.Error(err)
	}
	err = CreateUser("useR321", "pass123", "")
	if err != nil {
		t.Error(err)
	}
	err = CreateUser("user321", "pass123", "")
	if err != nil {
		// t.Error(err)
	}
	// PASS, duplicate username will return an error
}
