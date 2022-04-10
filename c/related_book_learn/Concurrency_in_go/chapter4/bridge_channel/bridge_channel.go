/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/21 23:45 8月
 **/
package main

import "fmt"

func main() {
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

	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		// 这个 channel 将返回所有桥接返回的结果
		valStream := make(chan interface{})
		go func() {  // 这个循环负责从 chanStream 中提取 channel 并将其提供给嵌套循环来使用
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				// 个人加：这里是个同步，依照了chanStream传来的顺序
				// 该循环负责读取已经给出的 channel 中的值，并将这些值重复到 valStream中。
				// 当我们当前正在循环的流关闭时，我们从执行从此 channel 读取的循环中跳出，并继续
				// 下一次迭代，选择要读取的channel。这为我们提供了一个不间断的结果值的流。
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	genVals := func() <-chan <-chan interface {} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i ++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}
	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
