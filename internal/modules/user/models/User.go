package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        string  `gorm:"varchar:191"`
	Email       string  `gorm:"varchar:191"`
	OTP         *string `gorm:"varchar:191"`
	Package     *uint8
	Role        string
	Password    string
	ActivatedAt *time.Time
}
