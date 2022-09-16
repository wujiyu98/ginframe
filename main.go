package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint = 19
	var b int = 2
	c := float64(a) / float64(b)
	fmt.Print(math.Ceil(c))

}
