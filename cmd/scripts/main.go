package main

import (
	"fmt"

	_ "github.com/wujiyu98/ginframe/config"
	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/model"
)

var db = database.DB

func Add() {
	seo := model.Seo{Name: "home", Pathname: "/", Meta: model.Meta{Title: "t1", Keywords: "ss,dd", Description: "faf"}}

	db.Create(&seo)

}

func Find() {
	var seo model.Seo

	db.First(&seo)
	fmt.Println(seo.Meta.Description)

}
func main() {
	Find()

}
