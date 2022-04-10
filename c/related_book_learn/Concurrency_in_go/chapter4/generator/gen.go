/**
 * @project let-sGo
 * @Author 27
 * @Description
 * @Date 2021/8/18 02:36 8月
 **/
package generator

func repeat(
	done <-chan interface{},
	values ...interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				//fmt.Println(v)
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

func take(
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
			case takeStream <- <-valueStream:  // 这里书上错了！！！！！
			}
		}
	}()
	return takeStream
}

func toString(
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