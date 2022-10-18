package main

import (
	"fmt"

	_ "github.com/wujiyu98/ginframe/config"
	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/entity"
)

var db = database.DB

func addData() {
	languages := []entity.Language{
		{Name: "English", Code: "en", Domain: "en"},
	}
	db.Create(&languages)

	catelogs := []entity.ArticleCategory{
		{Name: "company", Pathname: "company", ParentID: 0},
	}
	db.Create(&catelogs)

	var articles []entity.Article

	for i := 0; i < 10; i++ {

		articles = append(articles, entity.Article{Title: fmt.Sprint("at", i), LanguageID: 1, ArticleCategoryID: 1, Pathname: fmt.Sprint("at", i)})

	}
	db.Create(&articles)

}

func AutoMigrate() {
	// db.AutoMigrate(&entity.Meta{}, &entity.ArticleCategory{}, &entity.Address{}, &entity.Article{}, &entity.Banner{}, &entity.Category{}, &entity.Manufacturer{}, &entity.CategoryManufacturer{}, &entity.Language{}, &entity.Message{}, &entity.Enquiry{}, &entity.EnquiryProduct{}, &entity.Order{}, &entity.OrderProduct{}, &entity.Product{}, &entity.Attribute{}, &entity.ProductAttribute{}, &entity.User{}, &entity.Seo{}, &entity.SiteInfo{})

}

func AddArticles() {
	var rows []entity.Article
	for i := 0; i < 1000; i++ {
		row := entity.Article{
			Title:             fmt.Sprint("title", i),
			Pathname:          fmt.Sprint("title", i),
			ArticleCategoryID: 1,
			LanguageID:        1,
		}
		rows = append(rows, row)

	}
	db.Create(&rows)

}

func main() {
	AddArticles()
}
