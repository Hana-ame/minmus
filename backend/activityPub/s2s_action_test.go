package activityPub

import (
	"fmt"
	"testing"
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
