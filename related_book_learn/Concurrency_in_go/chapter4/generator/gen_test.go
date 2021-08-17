/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/18 02:33 8月
 **/
package generator

import "testing"

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N)) {

	}
}

func BenchmarkTyped(b *testing.B) {
	repeatT := func(
		done <-chan interface{},
		values ...string,
	) <-chan string {
		valueStream := make(chan string)
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

	takeT := func(
		done <-chan interface{},
		valueStream <-chan string,
		num int,
	) <-chan string {
		takeStream := make(chan string)
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

	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range takeT(done, repeatT(done, "a"), b.N) {

	}
}
