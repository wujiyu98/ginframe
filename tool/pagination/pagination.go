package pagination

import (
	"fmt"
	"html/template"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

func Default(page uint, count int64) *Paginate {
	p := &Paginate{
		Size:  10,
		Page:  page,
		Count: count,
		Slot:  3,
		Sort:  "id desc",
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
	p.Count = req.Count
	p.Size = req.Size
	if req.Sort == "" {
		req.Sort = "id-desc"
	}
	p.Sort = req.Sort
	p.Path = fmt.Sprint(ctx.Request.URL.Path)
	p.setOffset()
	return &p

}

type PaginateReq struct {
	Count int64  `form:"count" json:"count" `
	Page  uint   `form:"page" json:"page"`
	Size  uint   `form:"size" json:"size"`
	Sort  string `form:"sort" json:"sort"`
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
	Sort           string      `json:"sort"`
	Offset         int         `json:"offset"`
	PageCount      uint        `json:"page_count"`
	Data           interface{} `json:"data"`
}

//初始3个插槽即1...345...9 ，345这三个，如果需要多的可以添加1，3，5，7，9等数字
func (p *Paginate) SetSlot(slot uint) {
	p.Slot = slot
}

func (p *Paginate) AddKey(key string, value string) {
	if strings.Contains(p.Path, "?") {
		p.Path = p.Path + "&" + fmt.Sprintf("%s=%s", key, value)
	} else {
		p.Path = p.Path + "?" + fmt.Sprintf("%s=%s", key, value)
	}
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
				for i := 1; i <= int(p.Slot)+1; i++ {
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

func (p *Paginate) setPath() {

	p.FirstPageUrl = fmt.Sprint(p.Path)
	format := `&%spage=%d&size=%d&count=%d&sort=%s`

	if !strings.Contains(p.Path, "?") {
		format = `?%spage=%d&size=%d&count=%d&sort=%s`
	}

	p.LastPageUrl = template.HTMLEscapeString(fmt.Sprintf(format, p.Path, p.PageCount, p.Size, p.Count, p.Sort))
	if p.Page != 1 {
		p.PrevPageUrl = template.HTMLEscapeString(fmt.Sprintf(format, p.Path, p.Page-1, p.Size, p.Count, p.Sort))
	}
	if p.Page != p.PageCount {
		p.NextPageUrl = template.HTMLEscapeString(fmt.Sprintf(format, p.Path, p.Page+1, p.Size, p.Count, p.Sort))
	}

}

func (p *Paginate) BsPage() template.HTML {
	var html, navH, navF, prev, next, items string
	lists := p.GetList()
	p.setPath()
	navH = `<nav aria-label="pagination"> <ul class="pagination my-3">`
	navF = `</ul> </nav>`
	if p.PrevPageUrl == "" {
		prev = `<li class="page-item disabled"> <a class="page-link disabled" href="#" aria-label="Previous"> <span aria-hidden="true">&laquo;</span> </a> </li>`
	} else {
		prev = fmt.Sprintf(`<li class="page-item"> <a class="page-link" href="%s" aria-label="Previous"> <span aria-hidden="true">&laquo;</span> </a> </li>`, p.PrevPageUrl)
	}
	if p.NextPageUrl == "" {
		next = `<li class="page-item disabled"> <a class="page-link" href="#" aria-label="Next"> <span aria-hidden="true">&raquo;</span> </a> </li>`
	} else {
		next = fmt.Sprintf(`<li class="page-item"> <a class="page-link" href="%s" aria-label="Next"> <span aria-hidden="true">&raquo;</span> </a> </li>`, p.NextPageUrl)
	}
	for _, list := range lists {
		var item, linkUrl string
		format := `%s&page=%s&size=%d&count=%d&sort=%s`

		if !strings.Contains(p.Path, "?") {
			format = `%s?page=%s&size=%d&count=%d&sort=%s`
		}
		linkUrl = template.HTMLEscapeString(fmt.Sprintf(format, p.Path, list, p.Size, p.Count, p.Sort))

		switch {
		case list == "...":
			item = fmt.Sprintf(`<li class="page-item disabled"><a class="page-link" href="#">%s</a></li>`, list)
		case list == fmt.Sprint(p.Page):
			item = fmt.Sprintf(`<li class="page-item active" aria-current="page"><a class="page-link" href="%s">%s</a></li>`, linkUrl, list)
		default:
			item = fmt.Sprintf(`<li class="page-item"><a class="page-link" href="%s">%s</a></li>`, linkUrl, list)
		}
		items += item
	}
	if p.Count != 0 {
		html = navH + prev + items + next + navF
	}
	return template.HTML(html)

}
