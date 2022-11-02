package api

import (
	"html/template"

	rice "github.com/GeertJohan/go.rice"
	gintemplate "github.com/foolin/gin-template"
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"

	c "github.com/wujiyu98/ginframe/app/api/internal/controller"
)

func Init(e *gin.Engine) {

	mw := gintemplate.Middleware(gorice.NewWithConfig(rice.MustFindBox("../../views/backend"), gintemplate.TemplateConfig{
		Root:         "views/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: false,
		Delims:       gintemplate.Delims{Left: "{{", Right: "}}"},
	}))

	r := e.Group("/api", mw)

	r.GET("/hello", func(ctx *gin.Context) {
		gintemplate.HTML(ctx, 200, "index", ctx.Keys)
	})
	r.POST("/message", c.Message)
	r.POST("/enquiry", c.Enquity)

}
