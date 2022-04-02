//函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，
//使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，
//避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		_, err := fmt.Fprintf(os.Stderr, "Usage: %s url\n", os.Args[0])
		if err != nil {
			return
		}
		os.Exit(1)
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	nbytes, errCopy := io.Copy(os.Stdout, resp.Body)
	if errCopy != nil {
		panic(errCopy)
	}
	errClose := resp.Body.Close()
	if errClose != nil {
		return
	}
	fmt.Printf("%7d  %s\n", nbytes, url)
}
