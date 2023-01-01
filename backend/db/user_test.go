package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/hana-ame/minmus/backend/utils"
)

func TestCheckwhy(t *testing.T) {
	InitDB()

	var user User
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		fmt.Println("!!!", err)
	}

	if err := tx.First(&user).Error; err != nil {
		fmt.Println("!!!", err)
		tx.Rollback()
	}
	fmt.Println(user)

}
func TestCheck(t *testing.T) {
	InitDB()

	user := &User{
		Username: "test",
	}
	f, e := CheckUser(user)
	fmt.Println(f, e)
	// fmt.Println(f, e, user)

	user = &User{
		Username: "username123",
	}
	f, e = CheckUser(user)
	fmt.Println(f, e)
	// fmt.Println(f, e, user)

}

func TestCreat(t *testing.T) {

	InitDB()

	fmt.Println(db)

	privateKey := utils.GenerateKey()
	privateKeyPem := string(utils.MarshalPrivateKey(privateKey))
	publicKeyPem := string(utils.MarshalPublicKey(&privateKey.PublicKey))

	fmt.Println(privateKeyPem)
	fmt.Println(publicKeyPem)

	user := &User{
		ID:           utils.GetTS(),
		CreatedAt:    time.Now(),
		Username:     "username123",
		Email:        "sb@a.b",
		PasswordHash: "password123",
		// PasswordHash:  utils.Sha256String("password123"),
		PublicKeyPem:  publicKeyPem,
		PrivateKeyPem: privateKeyPem,
	}

	fmt.Println(user)

	// err := CreateUser(user)
	// if err != nil {
	// 	panic(err)
	// }

	db.Create(user)

}
