// 练习 3.13： 编写KB、MB的常量声明，然后扩展到YB。

package main

import (
	"fmt"
)

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("KB:", KB, "KB")
	fmt.Println("MB:", MB/KB, "KB")
	fmt.Println("GB:", GB/MB, "MB")
	fmt.Println("TB:", TB/GB, "GB")
	fmt.Println("PB:", PB/TB, "TB")
	fmt.Println("EB:", EB/PB, "PB")
	fmt.Println("ZB:", ZB/EB, "EB")
	fmt.Println("ZB:", YB/ZB, "ZB")
}
