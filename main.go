package main

import (
	"fmt"
	"time"

	_ "github.com/wujiyu98/ginframe/boot"
	"golang.org/x/time/rate"
	// "github.com/wujiyu98/ginframe/router"
)

func myLimit() {

	limit := rate.NewLimiter(0.083, 5)

	for i := 0; i < 70; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i + 1)

		fmt.Println(limit.Allow())

	}

}

func main() {
	m := map[string]string{"hello": "d", "mane": "sdfa"}
	for k := range m {
		delete(m, k)
	}
	m["age"] = "111"

	fmt.Println(len(m))
}
