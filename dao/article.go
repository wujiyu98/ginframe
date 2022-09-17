package dao

import (
	"fmt"

	"github.com/wujiyu98/ginframe/model"
	"github.com/wujiyu98/ginframe/pkg/pagination"
)

var Article = articleDao{}

type articleDao struct {
}

func (articleDao) All() (rows []model.Article) {
	db.Find(&rows)
	return
}

func (articleDao) Pagination(p *pagination.Paginate, query interface{}, args ...interface{}) (rows []model.Article) {
	var count int64
	if p.Count == 0 {
		if err := db.Model(&Article).Where(query, args...).Count(&count).Error; err != nil {
			fmt.Print(err)
		}
		p.SetCount(count)
	}

	db.Where(query, args...).Offset(p.Offset).Limit(int(p.Size)).Find(&rows)
	return
}
