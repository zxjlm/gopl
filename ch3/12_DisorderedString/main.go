// 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。

package main

import "fmt"

func main() {
	fmt.Println(isDisordered("abcd", "dbca"))    // true
	fmt.Println(isDisordered("abce", "dbca"))    // false
	fmt.Println(isDisordered("abcd123", "dbca")) // false
}

func isDisordered(s1, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}

	m := make(map[rune]int)
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}
