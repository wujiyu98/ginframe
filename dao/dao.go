package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/model"
	"github.com/wujiyu98/ginframe/tool/pagination"
	"gorm.io/gorm"
)

func New() *dao {
	return &dao{DB: database.DB}
}

type dao struct {
	*gorm.DB
}

func (d dao) pagination(p *pagination.Paginate, tx *gorm.DB, rows interface{}) {
	var count int64
	if p.Count == 0 {
		tx.Count(&count)
		p.SetCount(count)
	}
	tx.Order("desc id").Offset(p.Offset).Limit(int(p.Size)).Find(rows)

}

func (d dao) ProductPagination(ctx *gin.Context, size uint) (p *pagination.Paginate, rows []model.Article) {
	p = pagination.New(ctx, size)
	tx := d.Where("article_category_id", 1)
	d.pagination(p, tx, &rows)
	return p, rows

}
