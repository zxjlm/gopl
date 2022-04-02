//找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，
//并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。

//go run 10_FetchAllCache.go http://www.baidu.com http://www.sina.cn http://163.com http://www.qq.com http://www.weibo.cn

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"
)

const repeatTimes = 3

func main() {
	ch := make(chan string)
	urls := os.Args[1:]
	for _, url := range urls {
		go multiFetch(url, ch) // start a goroutine
	}
	var respStrs []string
	for i := 0; i < len(urls)*repeatTimes; i++ {
		//fmt.Println(<-ch) // receive from channel ch
		respStrs = append(respStrs, <-ch)
	}

	sort.Strings(respStrs)
	for _, s := range respStrs {
		fmt.Println(s)
	}
}

func multiFetch(url string, cha chan<- string) {
	client := http.Client{}
	for i := 0; i < repeatTimes; i++ {
		start := time.Now()
		//resp, err := http.Get(url)
		resp, err := client.Get(url)
		if err != nil {
			cha <- err.Error()
			return
		}
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		errClose := resp.Body.Close()
		if errClose != nil {
			return
		} // don't leak resources
		secs := time.Since(start).Seconds()
		//fmt.Printf("%.2fs %7d\n\n", secs, nbytes)

		cha <- fmt.Sprintf("%-20s %.2fs %7d  %7d", url, secs, resp.StatusCode, nbytes)
	}
	client.CloseIdleConnections()
}
