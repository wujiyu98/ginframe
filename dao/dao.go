package dao

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/tools/pagination"
	"gorm.io/gorm"
)

func New() *dao {
	return &dao{DB: database.DB}
}

type dao struct {
	*gorm.DB
}

func (d dao) pagination(p *pagination.Paginator, tx *gorm.DB, rows interface{}) {

	var count int64
	if p.Total == 0 {
		tx.Count(&count)
		p.Total = uint(count)
	}
	if p.Sort == "" {
		tx.Offset(p.Offset()).Limit(int(p.Size)).Find(rows)
	} else {
		tx.Order(strings.Replace(p.Sort, "-", " ", 1)).Offset(p.Offset()).Limit(int(p.Size)).Find(rows)
	}

}

//以表名取数据分页，如果超出页数超出，取最后一页;
// rows 是模型数组，query如果""是取全部
func (d dao) Pagination(table string, ctx *gin.Context, size uint, rows interface{}, query interface{}, args ...interface{}) *pagination.Paginator {
	p := pagination.New(ctx.Request, size)
	tx := d.Table(table).Where(query, args...)
	d.pagination(p, tx, rows)
	return p

}
