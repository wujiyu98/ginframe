package utils

import (
	"math"

	"github.com/gin-gonic/gin"
)

func NewPaginate(ctx *gin.Context, size uint, slot uint) (p *Paginate) {
	var req PaginateReq
	ctx.ShouldBindQuery(&req)
	if req.Size == 0 {
		req.Size = size
	}
	if req.Size > 100 {
		req.Size = 100
	}
	if req.Page == 0 {
		req.Page = 1
	}
	p.Page = req.Page
	p.Size = req.Size
	p.Slot = slot
	return
}

func Pagination(size uint, slot uint, page uint, count int64) *Paginate {
	return &Paginate{Size: size, Slot: slot, Page: page, Count: count}
}

type PaginateReq struct {
	Count int64 `form:"count" json:"count" `
	Page  uint  `form:"page" json:"page"`
	Size  uint  `form:"size" json:"size"`
}

type Paginate struct {
	Count          int64       `form:"count" json:"count" `
	Page           uint        `form:"page" json:"page"`
	Path           string      `form:"pathname" json:"pathname"`
	CurrentPageUrl string      `json:"current_page_url"`
	FirstPageUrl   string      `json:"first_page_url"`
	LastPageUrl    string      `json:"last_page_url"`
	PrevPageUrl    string      `json:"prev_page_url"`
	NextPageUrl    string      `json:"next_page_url"`
	Size           uint        `form:"size" json:"size"`
	Slot           uint        `json:"slot"`
	PageCount      uint        `json:"page_count"`
	Data           interface{} `json:"data"`
}

func (p *Paginate) init() {
	p.PageCount = uint(math.Ceil(float64(p.Count) / float64(p.Size)))

}
