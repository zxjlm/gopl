// 重写reverse函数，使用数组指针代替slice。

package main

import "fmt"

func reverse(ptrs *[5]int) {
	for i := 0; i < len(ptrs)/2; i++ {
		end := len(ptrs) - i - 1
		ptrs[i], ptrs[end] = ptrs[end], ptrs[i]
	}
}

func main() {
	sli := [5]int{1, 2, 3, 4, 5}
	reverse(&sli)
	fmt.Println(sli)
}
