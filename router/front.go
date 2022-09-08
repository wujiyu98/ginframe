package router

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/wujiyu98/ginframe/controller"
)

func init() {
	// servers other static files
	staticBox := rice.MustFindBox("../static")
	engine.StaticFS("/static", staticBox.HTTPBox())

	//new template engine
	engine.HTMLRender = gorice.New(rice.MustFindBox("../views"))

	c := controller.Front
	r := engine.Group("/")
	r.GET("/", c.Index)

}
