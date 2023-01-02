package activityPub

import (
	"fmt"
	"testing"

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
