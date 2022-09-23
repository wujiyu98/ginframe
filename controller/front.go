package controller

import (
	"fmt"

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
	p, rows := d.ProductPagination(ctx, 2)
	rep.Articles = rows
	fmt.Println(p)

	ctx.JSON(200, rep)

}
