package controller

import (
	"github.com/gin-gonic/gin"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {

	ctx.HTML(200, "index", nil)
}
