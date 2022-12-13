package general

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

}

// ref: https://bitfieldconsulting.com/golang/type-parameters
func Unmarshal[T any](data []byte, obj *T) (*T, error) {
	err := json.Unmarshal(data, obj)
	return obj, err
}

func Marshal[T any](obj T) ([]byte, error) {
	b, err := json.Marshal(obj)
	return b, err
}
