package models

// user model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;"`
	Name      string    `validate:"minLen:3" json:"name" gorm:"type:string"`
	Password  string    `validate:"required|minLen:6" json:"password" gorm:"type:string;"`
	Email     string    `validate:"required|email" json:"email" gorm:"type:string;unique"`
	Role      string    `json:"role" gorm:"type:string;default:admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = uuid.New().String()
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}
