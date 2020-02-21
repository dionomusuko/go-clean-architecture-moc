package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"unique;not null" json:"first-name" binding:"required"`
	Email     string     `gorm:"unique;not null" json:"email" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

type Users []User
