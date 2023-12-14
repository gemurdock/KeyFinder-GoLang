package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" gorm:"unique_index;varchar(25);not null"`
	Password  string    `json:"password" gorm:"varchar(255);not null"`
	CreatedAt time.Time `json:"created" gorm:"type:timestamptz; not null"`
	DeletedAt time.Time `json:"deleted" gorm:"type:timestamptz"`
	LastLogin time.Time `json:"last_login" gorm:"type:timestamptz"`
}
