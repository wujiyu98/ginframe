package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/dao"
	"github.com/wujiyu98/ginframe/model"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Test(ctx *gin.Context) {

}

func (c frontController) Index(ctx *gin.Context) {
	var rows []model.Article
	d := dao.New()
	p := d.Pagination("articles", ctx, 10, &rows, "")
	p.Position = "start"
	ctx.Set("page", p.Html())

	ctx.HTML(200, "index", ctx.Keys)
}

func (c frontController) Manufacturers(ctx *gin.Context) {

}

func (c frontController) Manufacturer(ctx *gin.Context) {

}

func (c frontController) About(ctx *gin.Context) {

}

func (c frontController) Info(ctx *gin.Context) {

}

func (c frontController) News(ctx *gin.Context) {

}

func (c frontController) Article(ctx *gin.Context) {

}

func (c frontController) Categories(ctx *gin.Context) {

}

func (c frontController) Category(ctx *gin.Context) {

}

func (c frontController) Product(ctx *gin.Context) {

}

func (c frontController) Enquiry(ctx *gin.Context) {

}
func (c frontController) Contact(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"url":  ctx.Request.URL.String(),
		"uri":  ctx.Request.RequestURI,
		"path": ctx.Request.URL.Path,
	})
}

func (c frontController) PostEnquiry(ctx *gin.Context) {

}

func (c frontController) PostMessage(ctx *gin.Context) {

}
