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

func AddProduct() {
	var product = model.Product{
		Title:          "t2",
		CategoryID:     1,
		ManufacturerID: 2,
		Pathname:       "t2",
		ProductAttributes: []model.ProductAttribute{
			{AttributeID: 3, Text: "asfsaf"},
			{AttributeID: 2, Text: "gasda88"},
		},
	}

	db.Create(&product)

}

func FindProduct() {
	var prodcut model.Product

	db.Preload("ProductAttributes").First(&prodcut)
	fmt.Println(prodcut)

}

func main() {

	FindProduct()

}
