/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/21 01:30 8月
 **/
package main

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
}
