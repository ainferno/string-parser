package main

import (
	"fmt"
	"parser/parser"
)

func main() {
	res, _ := parser.Parse("0{ab}2{ab}")
	// res := strings.Repeat("a", 0)
	fmt.Println(res)
}
