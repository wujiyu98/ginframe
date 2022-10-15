package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/dao"
	"github.com/wujiyu98/ginframe/model"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {
	var seo model.Seo

	dao.New().First(&seo)
	ctx.Set("seo", seo)

	ctx.HTML(200, "index", ctx.Keys)

}
