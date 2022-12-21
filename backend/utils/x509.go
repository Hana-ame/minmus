package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/fatih/color"
)

var fakeSalt = "meiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiromeiro"

type loopReader struct {
	s []byte
	p int
}

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

func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pk.(*rsa.PublicKey), nil
}

func GenerateKey(salt string) *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(&loopReader{[]byte(fakeSalt), 0}, 512)
	if err != nil {
		color.Red(err.Error())
	}
	return privateKey
}

func MarshalPublicKey(privateKey *rsa.PrivateKey) []byte {
	key, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		color.Red(err.Error())
		return nil
	}
	return key

}
