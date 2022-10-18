package router

import (
	"html/template"

	rice "github.com/GeertJohan/go.rice"
	gintemplate "github.com/foolin/gin-template"
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/wujiyu98/ginframe/controller"
)

func init() {
	// servers other static files
	staticBox := rice.MustFindBox("../static")
	engine.StaticFS("/static", staticBox.HTTPBox())

	//new template engine
	// engine.HTMLRender = gorice.New(rice.MustFindBox("../views"))
	engine.HTMLRender = gorice.NewWithConfig(rice.MustFindBox("../views"), gintemplate.TemplateConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: true,
		Delims:       gintemplate.Delims{Left: "{{", Right: "}}"},
	})

	c := controller.Front
	r := engine.Group("/")

	r.GET("/", c.Index)
	r.GET("/contact", c.Contact)

}
