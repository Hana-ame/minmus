package utils

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	r := Sha256([]byte{1, 2, 3})
	fmt.Printf("%s\n", r)
	fmt.Printf("%x\n", r)
	// should %x then it is hashstring that len=64
}
