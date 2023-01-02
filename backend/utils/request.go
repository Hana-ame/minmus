package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// added accept: application/activity+json, application/ld+json
func GetWithAccept(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/activity+json")
	req.Header.Set("Accept", "application/ld+json")

	return http.DefaultClient.Do(req)
	// r, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
}

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
