package webfinger

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	resource := getResource("username")
	if resource == nil {
		t.Error("resource is nil")
	}
	fmt.Println(resource)

	data, err := json.Marshal(resource)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(data))
	//PASS
}
