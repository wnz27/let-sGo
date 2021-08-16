# 错误处理

在并发程序中，错误处理可能难以正确进行。我们花了很多时间思考我们的各种 stage 
如何共享信息和进行协调，我们忘记考虑它们如何优雅地处理错误的状态。
当Go语言避开了流行的错误异常模型时，它声明错误处理非常重要，并且在开发我们的程序时，
我们对待报错的路径，【应该给予和算法相同的关注度】。
本着这种精神，让我们来看看在处理多个并发进程时我们如何做到这一点。

思考错误处理时最根本的问题是，"谁应该负责处理错误？" 在某些时候，
程序需要停止将错误输出来，并且实际上对它做了些什么。这么做的目的是什么？

在并发进程中，这个问题变得更复杂一些。因为并发进程独立于其父进程或兄弟进程运行，
它可能很难判断当错误出现的时候该做什么才是正确的。
查看下面代码以查看此问题的示例：
```go
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
```
[【demo】](problem_demo/problem_demo.go)

输出：
```shell
Response: 200 OK
Get "https://badhost": dial tcp: lookup badhost: no such host
Response: 200 OK
```
在这里我们看到在这个问题上goroutine没有选择。他不能简单地吞下错误，
因此它只能选择明智的事情来做：它会打印错误并希望某些内容被关注。
不要把你的goroutine放在这个尴尬的位置。

建议你分情况来考虑：一般来说，你的并发进程应该把他们的错误发送到你的程序的另一部分，
它有你的程序状态的完整信息，并可以做出更明智的决定。
以下示例演示了此问题的正确解决方案：
```go
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
```
[【demo】](solve_err_demo/solve_err_demo.go)

输出:
```shell
Response: 200 OK
error: Get "https://badhost": dial tcp: lookup badhost: no such host
Response: 200 OK
```

这也就是说我们在checkStatus这个goroutine 中产生的所有结果的结果集都可以被传递给我们
的 "main goroutine" 来对各种可能出现的错误进行处理。

从更广泛的角度来讲，我们已经成功地将错误处理的担忧从我们的生产者 goroutine 中分离出来。
这是可取的，因为生成 goroutine 的goroutine（在当前例子中是 main goroutine） 具有
更多的关于正在运行的程序的上下文，并且可以做出关于如何处理错误的更明智的决定。

在前面例子中，我们只是将错误写入 stdio（标准输入输出），但我们可以做其他的事情。

然我们稍微修改我们的程序，以便在出现三个或更多错误时停止尝试检查状态：
```go
done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"a", "https://www.baidu.com", "b", "c", "d"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount ++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Respose: %v\n", result.Response.Status)
	}
```
[【demo】](stop_err/stop_err.go)

输出：
```shell
error: Get "a": unsupported protocol scheme ""
Respose: 200 OK
error: Get "b": unsupported protocol scheme ""
error: Get "c": unsupported protocol scheme ""
Too many errors, breaking!
```
你可以看到，因为错误是从checkStatus 返回的而不是在goroutine内部处理的，
错误处理遵循熟悉的Go语言模式。这是一个简单的例子，但不难想象，main goroutine
正在协调多个goroutine 的结果，并制定更复杂的规则来继续或取消子goroutine。
此外，这里的主要内容是，在构建从goroutine返回值时，应将错误视为一等公民。
如果你的 goroutine 可能产生错误，那么这些错误应该与你的结果类型紧密结合，
并且通过相同的通信线传递，就像常规的同步函数一样。
