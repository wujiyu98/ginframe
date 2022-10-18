package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/dao"
	"github.com/wujiyu98/ginframe/reponse"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {
	var rep reponse.Index
	d := dao.New()
	p := d.Pagination("articles", ctx, 10, &rep.Articles, "")

	ctx.Set("page", p.BsPage())

	ctx.HTML(200, "index", ctx.Keys)
}

func (c frontController) Contact(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"url":  ctx.Request.URL.String(),
		"uri":  ctx.Request.RequestURI,
		"path": ctx.Request.URL.Path,
	})
}
