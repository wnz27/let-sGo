/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/19 23:56 8月
 **/
package base_func

import "sync"

func RepeatFn(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func ToString(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func ToInt(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

func Take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: // 这里书上错了！！！！！
			}
		}
	}()
	return takeStream
}

// 还是很精妙
func PrimeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer -= 1
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}


// 还是很精妙
func PrimeFinder2(done <-chan interface{}, intStream <-chan int) <-chan int {
	primeStream := make(chan int)
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer -= 1
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}

// FanIn 我们这里采用标准的 done channel 来使我们的 goroutine 可以被关闭，
// 然后用一个可变的interface{} 切片的channel 来进行 扇入
func FanIn(
	done <-chan interface{},
	channels ...<-chan interface{},
) <-chan interface{} {
	// 我们用一个 wait group 来等所有channel 都被处理完
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	// 在这里我们创建一个函数，他在传递时将从 channel 中读取，并将读取的值传递到 multiplexedStream channel
	multiplex := func(c <-chan interface{}) { // 处理单个channel 的策略
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// 从所有的 channel 里取值， 往wg 中增加channel 的数量
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)  // 实际扇出
	}

	// 等待所有的读操作结束
	go func() {
		// 我们创建一个goroutine 来等待我们多路复用 的 所有channel 被耗尽，
		// 这样我们可以关闭 multiplexedStream channel
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}


