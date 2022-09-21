package reponse

import "github.com/wujiyu98/ginframe/model"

type Index struct {
	Articles []model.Article
	Name     string
}
