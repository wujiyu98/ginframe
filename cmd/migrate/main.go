package main

import (
	_ "github.com/wujiyu98/ginframe/config"
	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/entity"
)

var db = database.DB

func main() {
	db.AutoMigrate(&entity.Language{}, &entity.ArticleCategory{}, &entity.Article{})

}
