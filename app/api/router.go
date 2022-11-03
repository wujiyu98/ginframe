package api

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"

	c "github.com/wujiyu98/ginframe/app/api/internal/controller"
)

func Init(e *gin.Engine) {

	r := e.Group("/api")

	r.GET("/hello", func(ctx *gin.Context) {
		gintemplate.HTML(ctx, 200, "index", ctx.Keys)
	})
	r.POST("/message", c.Message)
	r.POST("/enquiry", c.Enquity)

}
