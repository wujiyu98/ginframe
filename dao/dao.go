package dao

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
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

func (d dao) pagination(p *pagination.Paginate, tx *gorm.DB, sort string, rows interface{}) {

	var count int64
	if p.Count == 0 {
		tx.Count(&count)
		p.SetCount(count)
	}
	err := tx.Order(sort).Offset(p.Offset).Limit(int(p.Size)).Find(rows)

	if err != nil {
		fmt.Println(err.Error)

	}

}

//以表名取数据分页，如果超出页数超出，取最后一页;
// rows 是模型数组，query如果""是取全部
func (d dao) Pagination(table string, ctx *gin.Context, size uint, rows interface{}, query interface{}, args ...interface{}) *pagination.Paginate {
	p := pagination.New(ctx, size)
	sort := ctx.Query("sort")
	sort = strings.Replace(sort, "-", " ", 1)
	tx := d.Table(table).Where(query, args...)
	d.pagination(p, tx, sort, rows)
	return p

}
