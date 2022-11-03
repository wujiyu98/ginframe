package router

import (
	"html/template"

	rice "github.com/GeertJohan/go.rice"
	gintemplate "github.com/foolin/gin-template"
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/wujiyu98/ginframe/app/api"
	"github.com/wujiyu98/ginframe/app/front"
)

func Run() {
	e := gin.Default()
	e.SetTrustedProxies(nil)
	e.Use(ginsession.New())
	staticBox := rice.MustFindBox("../static")
	e.StaticFS("/static", staticBox.HTTPBox())
	//new template engine
	e.HTMLRender = gorice.New(rice.MustFindBox("../views"))
	e.HTMLRender = gorice.NewWithConfig(rice.MustFindBox("../views/frontend"), gintemplate.TemplateConfig{
		Root:         "views/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: true,
		Delims:       gintemplate.Delims{Left: "{{", Right: "}}"},
	})

	//下面是需要导入项目的路由

	front.Init(e)
	api.Init(e)

	// 路由运行
	e.Run()

}
