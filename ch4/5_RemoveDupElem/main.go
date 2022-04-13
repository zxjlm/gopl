// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

package main

import "fmt"

func removeEmpty(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			copy(s[i-1:], s[i:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

func main() {
	s := []string{"foo", "foo", "", "bar", "", "", "baz"}
	fmt.Println(removeEmpty(s))
}
