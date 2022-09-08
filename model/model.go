package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostCategoryID uint
	Title          string
	Content        string
}

type PostCategory struct {
	gorm.Model
	Name     string
	ParentID uint
}
