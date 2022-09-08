package router

import "github.com/gin-gonic/gin"

var engine = gin.Default()

func Run() {
	engine.Run()
}
