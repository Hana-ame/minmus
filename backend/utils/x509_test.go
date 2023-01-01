package utils

import (
	"fmt"
	"testing"
)

func TestGenPem(t *testing.T) {
	pk := GenerateKey()
	fmt.Println(pk)
}

func TestGenPemAndSave(t *testing.T) {
	pk := GenerateKey()
	// fmt.Println(pk)
	pubK := MarshalPublicKey(pk)
	fmt.Println(string(pubK))
	// fmt.Println(pk)
}

func TestGenPemAndSavePrivate(t *testing.T) {
	pk := GenerateKey()
	// fmt.Println(pk)
	pubK := MarshalPrivateKey(pk)
	fmt.Println(string(pubK))
	// fmt.Println(pk)
}
