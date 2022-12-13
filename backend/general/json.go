package general

import (
	"encoding/json"
	"log"
)

// ref: https://bitfieldconsulting.com/golang/type-parameters
func Unmarshal[T any](data []byte, obj *T) *T {
	err := json.Unmarshal(data, obj)
	if err != nil {
		log.Println(err)
		return nil
	}
	return obj
}

func Marshal[T any](obj T) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
