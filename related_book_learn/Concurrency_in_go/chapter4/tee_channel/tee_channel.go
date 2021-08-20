/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/21 01:55 8月
 **/
package main

import "fmt"

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

func main() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i ++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:  // 太他吗容易错了！！！！！
				}
			}
		}()
		return takeStream
	}

	orDone := func(
		done,
		c <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {  // 这里有死循环，所以在有死循环的时候要注意看下有没有从channel 读值并检查 channel 状态
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false { // 如果c关闭了
						return
					}
					// 不关闭则在起一个复用
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				// 我们要使用out1 和 out2 的私有变量版本，所以我们会覆盖这些变量。
				var out1, out2 = out1, out2
				// 我们将使用一条 select 语句，以便不阻塞的写入 out1 和 out2
				// 为确保两者都写入，我们将执行select 语句的两次迭代：每个出站一个channel（没看懂翻译， 看代码理解是每次传入一个）
				for i := 0; i < 2; i ++ {
					select {
					case <-done:
					// 一旦我们写入了channel，我们将其副本设置为nil，以便进一步阻塞写入，而另一个可以继续
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}

	done2 := make(chan interface{})
	for i := range ToInt(done, take(done2, repeat(done2, 1, 3), 4)) {
		fmt.Println(i)
	}
}
