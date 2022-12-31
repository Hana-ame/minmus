package activityPub

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	as := dummyPerson("user")
	fmt.Println(as)
}
