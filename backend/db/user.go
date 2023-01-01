package db

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// me / me@other.site
	Username string
	//
	Email string
	// sha256 hashed password
	PasswordHash string
	// public key
	PublicKeyPem string
	// private key
	PrivateKeyPem string

	// user infos
	// prefered name
	PreferredUsername string
	// display name
	Name string
	// summary
	Summary string
	// URL to pic
	Icon string
	// URL to pic
	Image string
	// Marshaled json string
	TagString string

	// settings
	ManuallyApprovesFollowers bool
}

type UserAS struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// me / me@other.site
	Username string
	// public key
	PublicKeyPem string

	// user infos
	// prefered name
	PreferredUsername string
	// display name
	Name string
	// summary
	Summary string
	// URL to pic
	Icon string
	// URL to pic
	Image string
	// Marshaled json string
	TagString string

	// settings
	ManuallyApprovesFollowers bool
}

// MUST and ONLY have Username, which should never be changed
func CheckUser(user *User) (bool, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Where("username = ?", user.Username).Find(user).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error
}

func CreateUser(user *User) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func QueryUser(username string) (*User, error) {
	var user User
	db.First(&user, "username = ?", username) //
	return &user, nil
}
