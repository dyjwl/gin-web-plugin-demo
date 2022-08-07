package model

import (
	"fmt"
	"time"

	"github.com/dyjwl/gin-web-plugin-demo/pkg/auth"
)

// User represents a user restful resource. It is also used as gorm model.
type User struct {
	Status int `json:"status" gorm:"column:status" validate:"omitempty"`
	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`
	// Required: true
	Email     string    `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Phone     string    `json:"phone" gorm:"column:phone" validate:"omitempty"`
	IsAdmin   int       `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`
	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

// UserList is the whole list of all users which have been stored in stroage.
type UserList struct {
	Total int64   `json:"total"`
	Items []*User `json:"items"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) error {
	if err := auth.Compare(u.Password, pwd); err != nil {
		return fmt.Errorf("failed to compile password: %w", err)
	}

	return nil
}
