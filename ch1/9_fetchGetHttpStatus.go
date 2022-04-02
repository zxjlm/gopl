// 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
}
