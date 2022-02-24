package parser

import (
	"strings"
)

func getSubstr(str string) int {
	count := 1
	for i, char := range str[1:] {
		if char == '{' {
			count++
		}

		if char == '}' {
			count--
		}

		if count == 0 {
			return i + 1
		}
	}
	return 0
}

func Parse(str string) (string, int) {
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '}' {
			return res, i
		}
		if str[i] >= '0' && str[i] <= '9' {
			a := int(str[i] - '0')

			res2, b := Parse(str[i+2:])

			res = res + strings.Repeat(res2, a)
			i += b + 2
		} else {
			res = res + string(str[i])
		}
	}
	return res, 0
}
