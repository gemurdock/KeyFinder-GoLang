package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string         `json:"username" gorm:"unique_index;varchar(25);not null"`
	Password  string         `json:"password" gorm:"varchar(255);not null"`
	LastLogin time.Time      `json:"last_login" gorm:"type:timestamptz"`
	CreatedAt time.Time      `json:"created" gorm:"type:timestamptz; not null"`
	UpdatedAt time.Time      `json:"updated" gorm:"type:timestamptz; not null"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}
