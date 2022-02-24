package main

import (
	"fmt"
	"parser/parser"
)

func main() {
	res := parser.Parse("2{}2{a}")
	// res := strings.Repeat("a", 0)
	fmt.Println(res)
}
