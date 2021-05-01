# 通道传输

## 示例：Telnet回音服务器——TCP服务器的基本结构

Telnet协议是TCP/IP协议族中的一种。
它允许用户(Telnet客户端)通过一个协商过程与一个远程设备进行通信。

本例将使用一部分Telnet协议与服务器进行通信。

服务器的网络库为了完整展示自己的代码实现了完整的收发过程，一般比较倾向于使用发送任意包返回原数据的逻辑。
这个过程类似于对着大山高喊，大山把你的声音原样返回的过程。也就是回音（Echo）。

本节使用Go语言中的Socket、goroutine和通道编写一个简单的Telnet协议的回音服务器。

回音服务器的代码分为4个部分，分别是:
1. 接受连接 
2. 会话处理
3. Telnet命令处理
4. 程序入口

#### 1．接受连接
回音服务器能同时服务于多个连接。要接受连接就需要先创建侦听器，侦听器需要一个侦听地址和协议类型。
- 主机IP：一般为一个IP地址或者域名，127.0.0.1表示本机地址。
- 端口号：16位无符号整型值，一共有65536个有效端口号。 
> 通过地址和协议名创建侦听器后，可以使用侦听器响应客户端连接。
响应连接是一个不断循环的过程，就像到银行办理业务时，一般是排队处理，前一个人办理完后，轮到下一个人办理。
![](../img/socket_handle.png)
> **Go语言中可以根据实际会话数量创建多个goroutine，并自动的调度它们的处理。**

telnet 服务器处理:
1．接受连接
```
func server(address string, exitChan chan int){
	// 根据给定地址进行侦听
	l, err := net.Listen("tcp", address)

	// 如果侦听发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	// 打印侦听地址，表示侦听成功
	fmt.Println("listen: ", address)

	// 延迟关闭侦听器
	defer l.Close()

	// 侦听循环
	for {
		// 新链接没有来的时候 Accept是阻塞的
		conn, err := l.Accept()

		// 发生任何的侦听错误，打印错误并退出服务器
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// 根据链接开启会话，这个过程需要并行执行
		go handleSession(conn, exitChan)
	}
}
``` 
2．会话处理
> 每个连接的会话就是一个接收数据的循环。当没有数据时，调用reader.ReadString会发生阻塞，
等待数据的到来。 一旦数据到来，就可以进行各种逻辑处理。
![](../img/telnet_handle_string.png)

回音服务器需要将收到的有效数据通过Socket发送回去。
```
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("session start")

	// 创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)

	// 接收数据的循环
	for {
		// 读取字符串，碰到回车返回
		str, err := reader.ReadString('\n')

		// 数据读取正确
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)

			// 处理Telnet 指令
			if !processTelnetCommend(str, exitChan){
				conn.Close()
				break
			}

			// Echo 逻辑 发送什么数据，原样返回
			conn.Write([]byte(str + "\r\n"))
		} else {
			// 发生错误
			fmt.Println("session closed")
			conn.Close()
			break
		}
	}
}
```
3．Telnet命令处理
Telnet是一种协议。在操作系统中可以在命令行使用Telnet命令发起TCP连接。
我们一般用Telnet来连接TCP服务器，键盘输入一行字符回车后，即被发送到服务器上。
在下例中，定义了以下两个特殊控制指令，用以实现一些功能：
- 输入“@close”退出当前连接会话。
- 输入“@shutdown”终止服务器运行。
```
func processTelnetCommend(str string, exitChan chan int) bool {
	// @close 指令表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("session closed")
		// 告诉外部需要断开连接
		return false

		// @shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown")

		// 往通道中写入0， 阻塞等待接收方处理
		exitChan <- 0

		// 告诉外部需要断开连接
		return false
	}

	// 打印输入的字符串
	fmt.Println("inner handle", str)
	return true
}

```
4．程序入口
Telnet 回音处理主流程
```
exitChan := make(chan int)

	// 将服务器并发运行
	go server("127.0.0.1:7001", exitChan)

	// 通道阻塞，等待接收返回值
	code := <- exitChan

	// 标记程序返回值并退出
	os.Exit(code)
```
todo: 事实上运行不动
命令不起作用






