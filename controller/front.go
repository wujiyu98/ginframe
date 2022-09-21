package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/dao"
	"github.com/wujiyu98/ginframe/model"
	"github.com/wujiyu98/ginframe/reponse"
	"github.com/wujiyu98/ginframe/tool/pagination"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {
	var rep reponse.Index
	d := dao.New()
	p := pagination.New(ctx, 2)
	d.Pagination(p, &model.Article{}, &rep.Articles, "article_category_id", 1)
	ctx.JSON(200, rep.Articles)

}
