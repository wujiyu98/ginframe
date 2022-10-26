package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["name"] = "wuj"
	for k, v := range m {
		fmt.Println(v, k)

	}

}
