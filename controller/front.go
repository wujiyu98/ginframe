package controller

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
)

var Front = frontController{}

type frontController struct {
}

func (c frontController) Index(ctx *gin.Context) {
	ctx.HTML(200, "index", nil)

}

func (c frontController) Upload(ctx *gin.Context) {
	s, _ := ctx.FormFile("file")

	tmp, _ := s.Open()
	defer tmp.Close()

	r := csv.NewReader(tmp)

	recodes, err := r.ReadAll()
	if err != nil {
		ctx.String(403, err.Error())
		return
	}

	ctx.JSON(200, recodes)
}
