package entity

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name   string `gorm:"size:60;not null"`
	Domain string `gorm:"size:120;not null"`
	Code   string `gorm:"size:10;not null"`
	Image  string `gorm:"size:255;default:''"`
}

type Meta struct {
	MetaTitle       string `gorm:"size:255;default:''"`
	MetaKeywords    string `gorm:"size:120;default:''"`
	MetaDescription string `gorm:"size:255;default:''"`
}

type Seo struct {
	gorm.Model
	LanguageID uint
	Name       string
	Pathname   string
	Meta
}

type SiteInfo struct {
	gorm.Model
	LanguageID  uint
	SiteName    string `gorm:"size:255;default:''"`
	Company     string `gorm:"size:255;default:''"`
	Contact     string `gorm:"size:255;default:''"`
	Phone       string `gorm:"size:20;default:''"`
	Phone2      string `gorm:"size:20;default:''"`
	MobilePhone string `gorm:"size:20;default:''"`
	Skype       string `gorm:"size:60;default:''"`
	QQ          string `gorm:"size:20;default:''"`
	Whatsapp    string `gorm:"size:20;default:''"`
	Address     string `gorm:"size:255;default:''"`
}

type ArticleCategory struct {
	gorm.Model
	Name      string `gorm:"size:255;not null"`
	Pathname  string `gorm:"size:255;not null;unique"`
	ParentID  uint   `gorm:"not null;default:0"`
	SortOrder byte   `gorm:"default:0"`
	Image     string `gorm:"size:255;default:''"`
	Summary   string `gorm:"type:text;"`
	Meta      Meta
}

type Article struct {
	gorm.Model
	LanguageID        uint
	ArticleCategoryID uint   `gorm:"not null"`
	Title             string `gorm:"size:255;not null;unique"`
	Pathname          string `gorm:"size:255;not null;unique"`
	SortOrder         uint   `gorm:"default:0"`
	Showed            byte   `gorm:"default:1"`
	Summary           string `gorm:"type:text;"`
	Image             string `gorm:"size:255;default:''"`
	Author            string `gorm:"size:60;default:''"`
	Content           string `gorm:"type:longtext"`
	Meta
}

type Category struct {
	gorm.Model
	Name      string `gorm:"size:255;not null"`
	Pathname  string `gorm:"size:255;not null;unique"`
	ParentID  uint   `gorm:"not null;default:0"`
	SortOrder byte   `gorm:"default:0"`
	Qty       uint   `gorm:"default:0"`
	Image     string `gorm:"size:255;default:''"`
	Summary   string `gorm:"type:text;"`
	Meta
}

type Manufacturer struct {
	gorm.Model
	Name      string `gorm:"size:255;not null"`
	Pathname  string `gorm:"size:255;not null;unique"`
	SortOrder byte   `gorm:"default:0"`
	Qty       uint   `gorm:"default:0"`
	Image     string `gorm:"size:255;default:''"`
	Summary   string `gorm:"type:text;"`
	Meta
}

type Product struct {
	gorm.Model
	LanguageID uint
	CategoryID uint     `gorm:"not null"`
	Title      string   `gorm:"size:255;not null;unique"`
	Pathname   string   `gorm:"size:255;not null;unique"`
	SortOrder  uint     `gorm:"default:0"`
	Showed     byte     `gorm:"default:1"`
	Summary    string   `gorm:"type:text;"`
	Image      string   `gorm:"size:255;default:''"`
	Images     []string `gorm:"type:json;serializer:json"`
	Author     string   `gorm:"size:60;default:''"`
	Content    string   `gorm:"type:longtext"`
	Stock      uint
	Hot        byte                 `gorm:"default:0"`
	New        byte                 `gorm:"default:0"`
	Special    byte                 `gorm:"default:0"`
	Price      float64              `gorm:"type:decimal(10,4)"`
	Prices     []map[string]float64 `gorm:"type:json;serializer:json"`
	Meta
}
