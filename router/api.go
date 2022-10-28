package router

import "github.com/wujiyu98/ginframe/controller"

func init() {
	c := controller.Api
	r := engine.Group("/api")
	r.POST("/message", c.Message)
	r.POST("/enquiry", c.Enquity)

}
