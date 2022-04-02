//修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
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
}
