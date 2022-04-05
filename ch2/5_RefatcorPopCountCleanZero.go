package main

import "fmt"

func PopCountZero(x uint64) int {
	var pcTmp int
	for x != 0 {
		// 将最低位的1置0
		x = x & (x - 1)
		pcTmp++
	}
	return pcTmp
}

func main() {
	fmt.Println(PopCountZero(18446744073709551615))
}
