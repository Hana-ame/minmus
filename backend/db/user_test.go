package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/hana-ame/minmus/backend/utils"
)

func TestUserAll(t *testing.T) {
	var err error
	var user *User
	InitDB()

	// create user
	user = &User{
		ID:           utils.GetTS(),
		CreatedAt:    time.Now(),
		Username:     "username123",
		Email:        "sb@a.b",
		PasswordHash: "password123",
		// PasswordHash:  utils.Sha256String("password123"),
		PublicKeyPem:  "",
		PrivateKeyPem: "",
	}
	err = CreateUser(user)
	if err != nil {
		t.Error(err)
	}

	user.Email = "not@allow.ed"
	user.PublicKeyPem = "PublicKeyPem"
	user.PrivateKeyPem = "PrivateKeyPem"
	err = CreateUser(user)
	if err == nil {
		t.Error("should not success")
	} else {
		fmt.Println(err)
	}

	user.Username = "username321"
	err = CreateUser(user)
	if err != nil {
		t.Error(err)
	}

	// check user
	var f bool
	user = &User{
		Username: "test",
	}
	f, err = CheckUser(user)
	// fmt.Println(f, err) // false recort not found
	if f {
		fmt.Println(f, err, user)
		t.Error("should not be found")
	}

	user = &User{
		Username: "username321",
	}
	f, err = CheckUser(user)
	// fmt.Println(f, err)
	if !f {
		fmt.Println(f, err, user)
		t.Error("should not be found")
	}

	// test delete
	user = &User{
		Username:     "username321",
		PasswordHash: "password123",
	}
	err = DeleteUser(user)
	if err != nil {
		t.Error(err)
	}

	// test update
	user = &User{
		Username:     "username123",
		PasswordHash: "password321", // this colum do not have any effect
	}
	_, err = QueryUser(user)
	if err != nil {
		t.Error(err)
	}

	user.PasswordHash = "new"
	err = UpdateUser(user)
	if err != nil {
		t.Error(err)
	}

	// test delete
	user = &User{
		Username:     "username123",
		PasswordHash: "password123",
	}
	err = DeleteUser(user)
	if err != nil {
		t.Error(err)
	}

}

func TestDelete(t *testing.T) {
	var err error
	var user *User
	InitDB()

	user = &User{
		Username:     "username321",
		PasswordHash: "password123",
	}
	err = DeleteUser(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	var err error
	var user *User
	InitDB()

	user = &User{
		Username:     "username123",
		PasswordHash: "password321", // this colum do not have any effect
	}
	_, err = QueryUser(user)
	if err != nil {
		t.Error(err)
	}

	user.PasswordHash = "new"
	err = UpdateUser(user)
	if err != nil {
		t.Error(err)
	}

}

func TestCheck(t *testing.T) {

	InitDB()

	user := &User{
		Username: "test",
	}
	f, e := CheckUser(user)
	fmt.Println(f, e)
	// false recort not found
	// fmt.Println(f, e, user)

	user = &User{
		Username: "username123",
	}
	f, e = CheckUser(user)
	fmt.Println(f, e)
	fmt.Println(f, e, user)
	if e == nil {
		t.Error(user)
	}
}

func TestCreate(t *testing.T) {

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
		Username:     "dummy",
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
