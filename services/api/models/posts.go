package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Slug        string `gorm:"unique"`
	Title       string `gorm:"unique"`
	Description *string
	Body        string
	ImageURL    *string
	UserID      uint
}

type PostService struct{}
