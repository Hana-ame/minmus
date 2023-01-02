package db

import (
	"time"
)

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	// me / me@other.site
	Username string `gorm:"primaryKey"`
	//
	Email string `gorm:"type:varchar(256)"`
	// sha256 hashed password
	PasswordHash string `gorm:"type:char(64)"`
	// private key
	PrivateKeyPem string `gorm:"type:varchar(2048)"`
	// endpoints
	Inbox       string `gorm:"type:varchar(512)"`
	Outbox      string `gorm:"type:varchar(512)"`
	Followers   string `gorm:"type:varchar(512)"`
	Following   string `gorm:"type:varchar(512)"`
	Featured    string `gorm:"type:varchar(512)"`
	SharedInbox string `gorm:"type:varchar(512)"`

	// public key
	PublicKeyID  string `gorm:"type:varchar(512)"`
	PublicKeyPem string `gorm:"type:varchar(2048)"`

	// user infos
	// prefered name, raw
	PreferredUsername string `gorm:"type:varchar(128)"`
	// display name
	Name string `gorm:"type:varchar(24)"`
	// summary
	Summary string `gorm:"type:text"`
	// URL to pic
	Icon string `gorm:"type:varchar(512)"`
	// URL to pic
	Image string `gorm:"type:varchar(512)"`
	// Marshaled json string
	TagString string `gorm:"type:varchar(512)"`

	// settings
	ManuallyApprovesFollowers bool
}

type UserAS struct {
	ID        int64
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

// MUST and ONLY have Username(which is PrimaryKey), which should never be changed
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

	if err := tx.First(user).Error; err != nil {
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

// should have only username(primary key). user returned is same as input.
func QueryUser(user *User) (*User, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return user, err
	}

	if err := tx.First(user).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	return user, tx.Commit().Error
}

// UpdateUser, use username to extinguish.
func UpdateUser(user *User) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// silent while username does not exist.
func DeleteUser(user *User) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
