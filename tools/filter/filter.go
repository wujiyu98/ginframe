package filter

import (
	"fmt"

	"gorm.io/gorm"
)

type Filter struct {
	Scopes []func(*gorm.DB) *gorm.DB
}

type Query struct {
	Key      string
	Operator string
	Arg      interface{}
}

func New() *Filter {
	return &Filter{}
}

func NewArray(rows []Query) {

}

//通用where查询
func (f *Filter) Query(query interface{}, args ...interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Where(query, args...)
	})
	return f

}

func (f *Filter) Equal(query string, arg interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Where(query, arg)
	})
	return f
}

//通用Not查询
func (f *Filter) Not(query interface{}, args ...interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Not(query, args...)
	})
	return f
}

func (f *Filter) Like(query string, arg interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Where(fmt.Sprintf(`%s like ?`, query), fmt.Sprintf(`%%%v%%`, arg))
	})
	return f
}

func (f *Filter) Gt(query string, arg interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {

		return d.Where(fmt.Sprintf(`%s > %v`, query, arg))

	})
	return f

}

func (f *Filter) Lt(query string, arg interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Where(fmt.Sprintf(`%s < %v`, query, arg))

	})
	return f
}

func (f *Filter) In(query string, arg interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Where(fmt.Sprintf(`%s in ?`, query), arg)

	})
	return f
}

func (f *Filter) Between(field string, arg1 interface{}, arg2 interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Where(fmt.Sprintf(`%s Between ? And ?`, field), arg2, arg2)
	})
	return f
}

func (f *Filter) Limit(limit int) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Limit(limit)

	})
	return f
}

func (f *Filter) Order(value interface{}) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Order(value)

	})
	return f
}

func (f *Filter) Offset(offset int) *Filter {
	f.Scopes = append(f.Scopes, func(d *gorm.DB) *gorm.DB {
		return d.Offset(offset)

	})
	return f
}

//operator 是运算符，其中^判断是不是排序
func (f *Filter) AddScope(query string, operator string, arg interface{}) *Filter {
	switch operator {
	case "=":
		f.Equal(query, arg)
	case "!=":
		f.Not(query, arg)
	case ">":
		f.Gt(query, arg)
	case "<":
		f.Lt(query, arg)
	case "^":
		sort := fmt.Sprintf(`%v %v`, query, arg)
		f.Order(sort)
	}
	return f

}
