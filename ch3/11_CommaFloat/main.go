// 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

package main

import (
	"strings"
)

func main() {
	println(comma("12345"))
	println(comma("12345.12345"))
	println(comma("-123456"))
	println(comma("-123456.12345"))
}

func comma(s string) string {
	var positive bool
	var result string

	if s != "" && s[0] == '-' {
		positive = false
		s = s[1:]
	} else {
		positive = true
	}

	if strings.Contains(s, ".") {
		splitRes := strings.Split(s, ".")
		left, right := splitRes[0], splitRes[1]
		result = commaIntegers(left) + "." + reverse(commaIntegers(reverse(right)))
	} else {
		result = commaIntegers(s)
	}

	if positive {
		return result
	} else {
		return "-" + result
	}
}

func commaIntegers(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
