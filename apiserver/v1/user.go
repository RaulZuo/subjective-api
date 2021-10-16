package v1

import (
	"gorm.io/gorm"
)

// User represents a user restful resource. It is also used as gorm model.
type User struct {
}

// UserList is the whole list of all users which have been stored in stroage.
type UserList struct {
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}

// Compare with the plain text password. Returns true if it's the same as the  encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	return
}

// BeforeCreate run before create database records.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return
}


