package controller

import (
	"github.com/gin-gonic/gin"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {

	ctx.String(200, "asdfa")
}
