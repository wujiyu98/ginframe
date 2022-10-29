package main

import (
	"fmt"

	_ "github.com/wujiyu98/ginframe/config"
	"github.com/wujiyu98/ginframe/tool/gaes"
)

func main() {
	key := "Jkasklgaslgl3ld4"
	s := gaes.EncryptString("", key)
	fmt.Println(s)

}
