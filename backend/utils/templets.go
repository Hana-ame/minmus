package utils

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	for k, v := range r.Header {
		color.Magenta(k)
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
