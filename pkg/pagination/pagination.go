package pagination

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
)

func Default(page uint, count int64) *Paginate {
	p := &Paginate{
		Size:  10,
		Page:  page,
		Count: count,
		Slot:  3,
	}
	p.setOffset()
	return p
}

func New(ctx *gin.Context, size uint, args ...uint) *Paginate {
	var req PaginateReq
	var p Paginate
	var slot uint = 3
	if len(args) > 0 {
		slot = args[0]
	}
	ctx.ShouldBindQuery(&req)
	if req.Size <= 0 {
		req.Size = size
	}
	if req.Size > 100 {
		req.Size = 100
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	p.Slot = slot
	p.Page = req.Page
	p.Size = req.Size

	p.setOffset()
	return &p

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
	Offset         int         `json:"offset"`
	PageCount      uint        `json:"page_count"`
	Data           interface{} `json:"data"`
}

func (p *Paginate) SetSlot(slot uint) {
	p.Slot = slot
}

func (p *Paginate) setOffset() {
	p.Offset = int(p.Page-1) * int(p.Size)
}

func (p *Paginate) SetCount(count int64) {
	p.Count = count
	p.setPageCount()
}
func (p *Paginate) setPageCount() {
	p.PageCount = uint(math.Ceil(float64(p.Count) / float64(p.Size)))
	if (p.Page > p.PageCount) && (p.PageCount > 0) {
		p.Page = p.PageCount
		p.setOffset()
	}
}

func (p *Paginate) GetList() (lists []string) {
	p.setPageCount()
	if p.PageCount > 0 {
		switch {
		case p.PageCount <= p.Slot+2:
			for i := 1; i < int(p.PageCount); i++ {
				lists = append(lists, fmt.Sprint(i))
			}
		case p.PageCount == p.Slot+3:
			if p.Page <= p.Slot {
				for i := 1; i <= int(p.Slot); i++ {
					lists = append(lists, fmt.Sprint(i))
				}
				lists = append(lists, "...")
				lists = append(lists, fmt.Sprint(p.PageCount))
			} else {
				lists = append(lists, "1")
				lists = append(lists, "...")
				for i := int(p.Slot); i <= int(p.PageCount); i++ {
					lists = append(lists, fmt.Sprint(i))
				}

			}
		default:
			switch {
			case p.Page <= p.Slot:
				for i := 1; i <= int(p.Slot); i++ {
					lists = append(lists, fmt.Sprint(i))
				}
				lists = append(lists, "...")
				lists = append(lists, fmt.Sprint(p.PageCount))
			case p.Page+p.Slot > p.PageCount:
				lists = append(lists, "1")
				lists = append(lists, "...")
				for i := int(p.PageCount - p.Slot); i <= int(p.PageCount); i++ {
					lists = append(lists, fmt.Sprint(i))
				}
			default:
				lists = append(lists, "1")
				lists = append(lists, "...")
				for i := int(p.Page - (p.Slot-1)/2); i <= int(p.Page+(p.Slot-1)/2); i++ {
					lists = append(lists, fmt.Sprint(i))
				}
				lists = append(lists, "...")
				lists = append(lists, fmt.Sprint(p.PageCount))
			}

		}

	}

	return

}
