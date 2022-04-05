package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var pcTmp byte
	for i := 0; i < 8; i++ {
		pcTmp += pc[byte(x>>(i*8))]
	}
	return int(pcTmp)
}

func main() {
	println(PopCount(18446744073709551615))
}
