package webcore

import (
	"fmt"

	"github.com/hana-ame/minmus/backend/db"
	"github.com/hana-ame/minmus/backend/utils"
)

func CreateUser(username string, password string, email string) error {
	if !utils.IsValidUsername(username) {
		return fmt.Errorf("invalid username")
	}
	pk := utils.GenerateKey()
	priPem := utils.MarshalPrivateKey(pk)
	pubPem := utils.MarshalPublicKey(&pk.PublicKey)
	passHash := utils.Sha256String(password)
	// TODO: 大小写？？？indexed能够去重，似乎
	user := &db.User{
		Username:      username,
		Email:         email,
		PasswordHash:  passHash,
		PrivateKeyPem: string(priPem),
		PublicKeyPem:  string(pubPem),
	}

	err := db.CreateUser(user)

	return err
}
