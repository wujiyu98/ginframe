package main

import (
	"fmt"

	"github.com/wujiyu98/ginframe/pkg/pagination"
)

func main() {

	p := pagination.Default(11, 100)

	fmt.Println(p.GetList())
	fmt.Print(p.Page)

}
