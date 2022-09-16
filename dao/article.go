package dao

import (
	"github.com/wujiyu98/ginframe/model"
)

var Article = articleDao{}

type articleDao struct {
}

func (articleDao) All() (rows []model.Article) {
	db.Find(&rows)
	return
}
