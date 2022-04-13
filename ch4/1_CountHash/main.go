// 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)

package main

import (
	"crypto/sha256"
	"fmt"
)

func countHash(c1, c2 [32]byte) {
	var count int
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			count += diffBitCount(c1[i] ^ c2[i])
		}
	}
	fmt.Printf("%d\n", count)
}

func diffBitCount(diff byte) int {
	count := 0
	for diff != 0 {
		count++
		diff &= diff - 1
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", c1, c2)
	countHash(c1, c2)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// 125
}
