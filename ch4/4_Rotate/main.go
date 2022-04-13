// 编写一个rotate函数，通过一次循环完成旋转。
// ex4.4 rotates a slice of ints by one position to the left.
// ps. 令人迷惑的翻译。。。

package main

import (
	"fmt"
)

func rotate(s []int, n int) {
	tmp := append(s, s[:n]...)
	copy(s, tmp[n:])
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	rotate(s, 2)
	fmt.Println(s) // [2 3 4 0 1]
}
