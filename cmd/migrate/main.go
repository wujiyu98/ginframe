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

func main() {
	db.AutoMigrate(&entity.Language{}, &entity.ArticleCategory{}, &entity.Article{})
	addData()
}
