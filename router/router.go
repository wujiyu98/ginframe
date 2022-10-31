package router

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

var engine = start()

func start() *gin.Engine {
	e := gin.Default()
	e.SetTrustedProxies(nil)
	// 开启session
	e.Use(ginsession.New())
	return e
}

func Run() {
	engine.Run()
}
