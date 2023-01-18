package dao

import (
	"fmt"

	"github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/tools/filter"
	"gorm.io/gorm"
)

func New() *Dao {

	return &Dao{DB: database.DB}
}

func Table(table string) *Dao {
	d := &Dao{DB: database.DB}
	d.DB = d.Table(table)
	return d

}

type Dao struct {
	*gorm.DB
}

func (d *Dao) Query(f *filter.Filter) *Dao {
	fmt.Println(f.Scopes)

	d.DB = d.Scopes(f.Scopes...)

	return d

}

//用maps 返回
func (d *Dao) All() (rows []map[string]interface{}) {
	d.DB.Find(&rows)
	return
}

func (d *Dao) FindAll() (rows interface{}) {
	d.DB.Find(&rows)
	return
}
