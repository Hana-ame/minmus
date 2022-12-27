package utils

import (
	"fmt"
	"testing"
)

func TestTS(t *testing.T) {
	// ts := &TS{}
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(GetTS())
		}
	}()
	for i := 0; i < 1000; i++ {
		fmt.Println(GetTS())
	}
}

func Test1(t *testing.T) {
	i := 1
	for j := 0; j < 66; j++ {
		i <<= 1
		fmt.Println(i, i<<16)
	}
}
