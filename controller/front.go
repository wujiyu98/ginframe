package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/dao"
	"github.com/wujiyu98/ginframe/pkg/pagination"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {
	p := pagination.New(ctx, 10)
	rows := dao.Article.Pagination(p, "article_category_id", 1)
	fmt.Print(rows)
	ctx.JSON(200, gin.H{"articles": rows, "lists": p.GetList()})
}
