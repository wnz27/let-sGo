/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/12 02:54 8月
 **/
package main

func main() {
	// 从nil channel 读
	//var dataStream chan interface{}
	//<- dataStream

	// 往nil channel 写
	//var dataStream chan interface{}
	//dataStream <- struct{}{}

	// 关闭 nil channel
	var dataStream chan interface{}
	close(dataStream)
}
