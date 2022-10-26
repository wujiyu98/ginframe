package pagination

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Paginator struct {
	Total     uint
	PageCount uint
	Size      uint
	Page      uint
	Query     map[string]string
	Path      string
	Sort      string
	From      string
	Slot      uint
	Data      interface{}
}

var DefaultConfig = Paginator{
	Page: 1,
	Slot: 5,
	Size: 10,
	Path: "/",
	Sort: "id-desc",
}

func New(r *http.Request, prePage uint) *Paginator {
	var p Paginator
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	total, _ := strconv.Atoi(q.Get("total"))
	size, _ := strconv.Atoi(q.Get("size"))
	sort := q.Get("sort")
	p.Page = uint(page)
	p.Total = uint(total)
	p.Size = uint(size)
	if p.Size == 0 {
		p.Size = prePage
	}
	p.Sort = sort
	p.validParams()
	return &p

}
func Default(total uint, size uint) *Paginator {
	p := Paginator{
		Total: total,
		Size:  size,
	}
	p.validParams()
	p.setPageCount()
	return &p
}

func (p *Paginator) SetPath(path string) {
	p.Path = path
}

func (p *Paginator) setPageCount() {
	p.PageCount = uint(math.Ceil(float64(p.Total) / float64(p.Size)))
	if p.Page > p.PageCount {
		p.Page = p.PageCount
	}

}

func (p *Paginator) Offset() int {
	return int((p.Page - 1) * p.Size)
}

func (p *Paginator) validParams() {
	p.Query = make(map[string]string)
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Slot == 0 {
		p.Slot = DefaultConfig.Slot
	}
	if p.Sort == "" {
		p.Sort = DefaultConfig.Sort
	}
	if p.Path == "" {
		p.Path = "/"

	}
}

func (p *Paginator) Paginate() (lists []string) {
	p.setPageCount()
	if p.Total > 0 {
		switch {
		case p.PageCount <= p.Slot:
			for i := 1; i <= int(p.PageCount); i++ {
				lists = append(lists, fmt.Sprint(i))
			}
		default:
			switch {
			case p.Page > p.PageCount-(p.Slot+1)/2:
				lists = append(lists, `1`)
				lists = append(lists, `...`)
				for i := int(p.PageCount - (p.Slot - 2)); i <= int(p.PageCount); i++ {
					lists = append(lists, fmt.Sprint(i))
				}

			case p.Page <= (p.Slot+1)/2:
				for i := 1; i <= int(p.Slot-1); i++ {
					lists = append(lists, fmt.Sprint(i))
				}
				lists = append(lists, `...`)
				lists = append(lists, fmt.Sprint(p.PageCount))

			default:
				lists = append(lists, `1`)
				lists = append(lists, `...`)
				for i := int(p.Page - (p.Slot-2)/2); i <= int(p.Page+(p.Slot-3)/2); i++ {
					lists = append(lists, fmt.Sprint(i))
				}
				lists = append(lists, `...`)
				lists = append(lists, fmt.Sprint(p.PageCount))

			}
		}

	}

	return

}

func (p *Paginator) PageItems() (items []map[string]string) {
	lists := p.Paginate()
	path := p.QueryUrl()
	for _, list := range lists {
		var url string
		if list != "..." {
			url = p.addUrlParam(path, "page", list)
			url = p.setComonParam(url)
		}
		m := make(map[string]string)
		m[list] = url
		items = append(items, m)
	}
	return
}

func (p *Paginator) AddQuery(key string, value string) {

	p.Query[key] = value

}

func (p *Paginator) AddQueries(items map[string]string) {

	for k, v := range items {
		p.Query[k] = v

	}

}

func (p *Paginator) addUrlParam(path string, k string, v string) string {
	if !strings.Contains(path, "?") {
		path += fmt.Sprint("?", k, "=", template.HTMLEscapeString(v))
	} else {
		path += fmt.Sprint("&", k, "=", template.HTMLEscapeString(v))
	}
	return path
}

func (p *Paginator) QueryUrl() string {
	path := p.Path
	for k, v := range p.Query {
		path = p.addUrlParam(path, k, v)
	}
	return path
}

func (p *Paginator) FirstPageUrl() string {
	path := p.QueryUrl()
	return path
}

func (p *Paginator) PreviousPageUrl() (url string) {
	if !(p.Page == 1) {
		path := p.QueryUrl()
		url = p.addUrlParam(path, "page", fmt.Sprint(p.Page-1))
		url = p.setComonParam(url)
	}
	return
}

func (p *Paginator) NextPageUrl() (url string) {
	if p.Page != p.PageCount {
		path := p.QueryUrl()
		url = p.addUrlParam(path, "page", fmt.Sprint(p.Page+1))
		url = p.setComonParam(url)
	}
	return
}

func (p *Paginator) setComonParam(path string) string {
	path = p.addUrlParam(path, "total", fmt.Sprint(p.Total))
	path = p.addUrlParam(path, "size", fmt.Sprint(p.Size))
	path = p.addUrlParam(path, "sort", p.Sort)
	return path

}

func (p *Paginator) LastPageUrl() string {
	path := p.QueryUrl()
	path = p.addUrlParam(path, "page", fmt.Sprint(p.PageCount))
	return p.setComonParam(path)
}

func (p *Paginator) Html() template.HTML {
	var html, navL, navR, prev, next, content string

	items := p.PageItems()
	if len(items) == 0 {
		return template.HTML(html)
	}
	navL = `<nav aria-label="pagination"> <ul class="pagination my-3">`
	navR = `</ul> </nav>`
	previousPageUrl := p.PreviousPageUrl()
	if previousPageUrl == "" {
		prev = `<li class="page-item disabled"> <a class="page-link disabled" href="#" aria-label="Previous"> <span aria-hidden="true">&laquo;</span> </a> </li>`
	} else {
		prev = fmt.Sprintf(`<li class="page-item"> <a class="page-link" href="%s" aria-label="Previous"> <span aria-hidden="true">&laquo;</span> </a> </li>`, previousPageUrl)
	}
	nextPageUrl := p.NextPageUrl()
	if nextPageUrl == "" {
		next = `<li class="page-item disabled"> <a class="page-link" href="#" aria-label="Next"> <span aria-hidden="true">&raquo;</span> </a> </li>`
	} else {
		next = fmt.Sprintf(`<li class="page-item"> <a class="page-link" href="%s" aria-label="Next"> <span aria-hidden="true">&raquo;</span> </a> </li>`, nextPageUrl)
	}
	currentpageStr := fmt.Sprint(p.Page)
	for _, v := range items {
		for k, v1 := range v {
			switch k {
			case `...`:
				content += fmt.Sprintf(`<li class="page-item disabled" aria-current="..."><a class="page-link">%s</a></li>`, k)
			case currentpageStr:
				if k == "1" {
					content += fmt.Sprintf(`<li class="page-item active" aria-current="page"><a class="page-link" href="%s">%s</a></li>`, p.FirstPageUrl(), k)
				} else {
					content += fmt.Sprintf(`<li class="page-item active" aria-current="page"><a class="page-link" href="%s">%s</a></li>`, v1, k)
				}
			default:
				if k == "1" {
					content += fmt.Sprintf(`<li class="page-item" aria-current="page"><a class="page-link" href="%s">%s</a></li>`, p.FirstPageUrl(), k)
				} else {
					content += fmt.Sprintf(`<li class="page-item" aria-current="page"><a class="page-link" href="%s">%s</a></li>`, v1, k)

				}

			}

		}
	}

	html = navL + prev + content + next + navR

	return template.HTML(html)

}
