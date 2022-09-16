package entity

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name     string `gorm:"size:60;not null"`
	Domain   string `gorm:"size:120;not null"`
	Code     string `gorm:"size:10;not null"`
	ImageSrc string `gorm:"size:255;default:''"`
}

type ArticleCategory struct {
	gorm.Model
	Name            string `gorm:"size:255;not null;unique"`
	Pathname        string `gorm:"size:255;not null;unique"`
	ParentID        uint   `gorm:"not null"`
	Sort            uint   `gorm:"default:0"`
	Summary         string
	MetaTitle       string `gorm:"size:255;not null"`
	MetaKeywords    string `gorm:"size:120;default:''"`
	MetaDescription string `gorm:"size:255;default:''"`
}

type Article struct {
	gorm.Model
	LanguageID        uint
	ArticleCategoryID uint   `gorm:"not null"`
	Title             string `gorm:"size:255;not null;unique"`
	Pathname          string `gorm:"size:255;not null;unique"`
	Sort              uint   `gorm:"not null;default:0"`
	Showed            uint   `gorm:"not null;default:1"`
	Summary           string `gorm:"size:1000"`
	ImageSrc          string
	Author            string
	Content           string `gorm:"type:longtext"`
	MetaTitle         string `gorm:"size:255;not null"`
	MetaKeywords      string `gorm:"size:120;default:''"`
	MetaDescription   string `gorm:"size:255;default:''"`
}
