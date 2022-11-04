package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/tools/throttle"
)

var IpLimit = throttle.NewIPRateLimiter(1, 5)

func Throttle(ctx *gin.Context) {
	limit := IpLimit.GetLimiter(ctx.ClientIP())
	if !limit.Allow() {
		ctx.AbortWithStatus(429)
	}

}
