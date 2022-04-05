package main

import "fmt"

func PopCountPosition(x uint64) uint64 {
	var pcTmp uint64
	for i := 0; i < 64; i++ {
		// 检测最后一位是否为1
		pcTmp += (x >> i) & 1
	}
	return pcTmp
}

func main() {
	fmt.Println(PopCountPosition(18446744073709551615))
	fmt.Println(PopCountPosition(72057594037927935))
}
