/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/16 23:46 8月
 **/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					// goroutine 在尽最大努力表现出错误，他还能做什么？它不能简单地将错误传回！
					// 它认为有多少错误才是太多呢？
					// 它是否需要继续请求
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://badhost", "https://www.baidu.com"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
