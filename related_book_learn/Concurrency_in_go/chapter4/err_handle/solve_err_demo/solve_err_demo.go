/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/17 00:04 8月
 **/
package main

import (
	"fmt"
	"net/http"
)

// Result 在这里，我们创建一个可以同时包含goroutine 中的循环迭代中的 * http.Response 以及可能出现的error的类型
type Result struct {
	Error error
	Response *http.Response
}

func main() {
	// 返回一个可读取的channel， 以检索循环迭代的结果。
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				// 在这里，我们创建一个Result 实例，并设置错误和响应字段
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:  // 将结果写入channel的地方。
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://badhost", "https://www.baidu.com"}
	for result := range checkStatus(done, urls...) {
		// 在我们的main goroutine 中，我们能够只能地处理由checkStatus 启动的goroutine中出现的错误，
		// 以及程序的全部上下文。
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
