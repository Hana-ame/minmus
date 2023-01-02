package webfinger

import "fmt"

func GetIdFromWebfinger(res Resource) (string, error) {
	for _, link := range res.Links {
		if link.Rel == "self" {
			return link.HRef, nil
		}
	}
	return "", fmt.Errorf("not webfinger, find no links with rel=self")
}
