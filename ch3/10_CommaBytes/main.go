// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

// 输入comma函数的参数是一个字符串。如果输入字符串的长度小于或等于3的话，则不需要插入逗分隔符。否则，
// comma函数将在最后三个字符前位置将字符串切割为两个两个子串并插入逗号分隔符，然后通过递归调用自身来出前面的子串。

package main

import "fmt"

func main() {
	fmt.Println(string(comma([]byte("123456789"))))
}

func comma(buffers []byte) []byte {
	result := make([]byte, 0)
	for idx, buffer := range buffers {
		if idx%3 == 0 {
			if idx == 0 {
				result = append(result, buffer)
				continue
			}
			result = append(result, ',')
		}
		result = append(result, buffer)
	}
	return result
}
