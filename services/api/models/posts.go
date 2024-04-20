package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          uint   `gorm:"primarykey"`
	Slug        string `gorm:"unique"`
	Title       string `gorm:"unique"`
	Description *string
	Body        string
	ImageURL    *string
	UserID      uint
}

type PostService struct {
	DB *gorm.DB
}

func (ps *PostService) Create(title string, description string, body string, user User) (*Post, error) {
	newPost := Post{
		Title:       title,
		Description: &description,
		Body:        body,
		UserID:      user.ID,
	}
	result := ps.DB.Create(&newPost)
	return &newPost, result.Error
}

func (ps *PostService) GetPostById(id uint) (*Post, error) {
	var post = Post{ID: id}
	result := ps.DB.First(&post)
	return &Post{}, result.Error
}

func (ps *PostService) GetAllPosts() (*[]Post, error) {
	posts := []Post{}
	ps.DB.Find(&posts)
	return &posts, nil
}
