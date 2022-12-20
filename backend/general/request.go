package general

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return []byte(resp.Status)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	//Convert the body to type string

	return body
}
