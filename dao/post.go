package dao

import (
	"github.com/wujiyu98/ginframe/model"
)

var Post = postDao{}

type postDao struct {
}

func (postDao) All() []model.Post {
	var rows []model.Post
	db.Find(&rows)
	return rows
}
