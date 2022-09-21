package dao

import (
	"fmt"

	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/tool/pagination"
	"gorm.io/gorm"
)

func New() *dao {
	return &dao{DB: database.DB}
}

type dao struct {
	*gorm.DB
}

func (d dao) Pagination(p *pagination.Paginate, m interface{}, rows interface{}, query interface{}, args ...interface{}) {
	var count int64
	if p.Count == 0 {
		if err := d.Model(m).Where(query, args...).Count(&count).Error; err != nil {
			fmt.Print(err)
		}
		p.SetCount(count)
	}
	d.Where(query, args...).Offset(p.Offset).Limit(int(p.Size)).Find(rows)

}
