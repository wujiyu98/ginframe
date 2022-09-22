package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/reponse"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {
	var rep reponse.Index

	ctx.JSON(200, rep.Articles)

}
