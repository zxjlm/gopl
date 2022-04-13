// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

package main

import "unicode/utf8"

func reverse(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
	return b
}

func main() {
	b := []byte("Hello, 世界")
	reverseUTF8(b)
	println(string(b))
}
