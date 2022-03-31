package main

import (
	"crypto/sha256"
	"fmt"
)

func CountHash(s string) int {
	count := len(sha256.Sum256([]byte(s)))
	return count
}

func main() {
	resX := CountHash("x")
	fmt.Println(resX)

	resY := CountHash("X")
	fmt.Println(resY)
}
