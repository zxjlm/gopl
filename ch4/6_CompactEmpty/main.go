// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

package main

import "unicode"

func compactSpace(s []byte) []byte {
	// use window to record the continuous space
	right := 0
	for left := 0; left < len(s); left++ {
		if unicode.IsSpace(rune(s[left])) {
			right = left + 1
			for right < len(s) {
				if unicode.IsSpace(rune(s[right])) {
					right++
				} else {
					break
				}
			}
			copy(s[left+1:], s[right:])
			s = s[:len(s)-(right-left-1)]
		}
	}
	return s
}

func main() {
	s := []byte("a b    c d   e  f   g")
	s = compactSpace(s)
	println(string(s))
}
