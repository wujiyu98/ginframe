package main

import (
	"fmt"
	"html/template"
	"net/url"
)

func main() {
	s := `"><script>alert(111);</script>"`

	fmt.Println(template.HTMLEscapeString(url.PathEscape(s)))

}
