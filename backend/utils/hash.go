package utils

import "crypto/sha256"

func Sha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func Sha256String(data string) string {
	return string(Sha256([]byte(data)))
}

func Sha256StringSalt(data string, salt string) string {
	return Sha256String(salt + data + salt)
}
