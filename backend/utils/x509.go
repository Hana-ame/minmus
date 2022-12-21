package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

// not used
// var fakeSalt = "meiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiro"

// not used
type loopReader struct {
	s []byte
	p int
}

// not used
func (r *loopReader) Read(b []byte) (n int, err error) {
	l := len(r.s)
	p := 0
	for p < len(b) {
		r.p %= l
		n := copy(b[p:], r.s[r.p:])
		r.p += n
		p += n
		fmt.Println(p, len(b))
	}
	return p, nil
}

// parse pem string to public key
func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pk.(*rsa.PublicKey), nil
}

// parse pem string to private key
func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	pk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pk.(*rsa.PrivateKey), nil
}

// create a new private key
func GenerateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		color.Red(err.Error())
	}
	return privateKey
}

// marshal public key to pem format
func MarshalPublicKey(privateKey *rsa.PrivateKey) []byte {
	bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		color.Red(err.Error())
		return nil
	}
	publicPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: bytes,
		},
	)
	return publicPem
}

// marshal private key to pem format
func MarshalPrivateKey(privateKey *rsa.PrivateKey) []byte {
	bytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		color.Red(err.Error())
		return nil
	}
	privatePem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: bytes,
		},
	)
	return privatePem
}

func ReadKeyFromFile(path string) (*rsa.PrivateKey, error) {
	return readKeyFromFile(path)
}

// read pem file from path, create file when it not exists
func readKeyFromFile(path string) (*rsa.PrivateKey, error) {
	if privateKeyPem, err := ioutil.ReadFile(path); err != nil {
		privateKey := GenerateKey()
		data := MarshalPrivateKey(privateKey)
		if data == nil {
			return nil, errors.New("data is nil")
		}
		err = os.WriteFile(path, data, 0644)
		// fmt.Println(err.Error()) // The system cannot find the file specified.
		if err != nil {
			return nil, err
		}
		return privateKey, nil
	} else {
		privateKey, err := ParsePrivateKey(string(privateKeyPem))
		if err != nil {
			color.Red(err.Error())
			return nil, err
		}
		return privateKey, nil
	}
}
