package controller

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {

}

func Index(ctx *gin.Context) {

	gintemplate.HTML(ctx, 200, "index", ctx.Keys)
}

func Manufacturers(ctx *gin.Context) {

}

func Manufacturer(ctx *gin.Context) {

}

func About(ctx *gin.Context) {

}

func Info(ctx *gin.Context) {

}

func NewsCategory(ctx *gin.Context) {

}

func News(ctx *gin.Context) {

}

func Categories(ctx *gin.Context) {

}

func Category(ctx *gin.Context) {

}

func Product(ctx *gin.Context) {

}

func Enquiry(ctx *gin.Context) {

}
func Contact(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"url":  ctx.Request.URL.String(),
		"uri":  ctx.Request.RequestURI,
		"path": ctx.Request.URL.Path,
	})
}
