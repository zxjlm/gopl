// 编写一个程序，默认打印标准输入的以SHA256哈希码，也可以通过命令行标准参数选择SHA384或SHA512哈希算法。
// go run main.go -t sha384 x

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashType = flag.String("t", "sha256", "hash type")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}
	//fmt.Println(args[0], *hashType)

	switch *hashType {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(args[0])))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(args[0])))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(args[0])))
	default:
		fmt.Println("invalid hash type")
	}

}
