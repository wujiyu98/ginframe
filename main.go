package main

import (
	"encoding/json"
	"fmt"
)

type Cart struct {
	ID       uint
	Quantity uint
}

func main() {
	// key := "fsafasfa#sgasf6f"
	// s := gaes.EncryptString(`[1,23,4,5,6]`, key)
	// fmt.Println(s)
	// o, _ := gaes.DecryptString(s, key)
	// fmt.Print(o)
	var list []Cart

	s := `[{"id":1,"quantity":131},{"id":3,"quantity":"131"}]`

	json.Unmarshal([]byte(s), &list)
	fmt.Println(list)
}
