// 修改dup2，出现重复的行时打印文件名称。
// go run 4_DupPlus.go test_datas/test1.txt test_datas/test2.txt test_datas/test3.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			if err != nil {
				return
			}
			continue
		}
		countLines(f, counts)
		errClose := f.Close()
		if errClose != nil {
			return
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("duplicate line: %s\t%s\n", f.Name(), input.Text())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
