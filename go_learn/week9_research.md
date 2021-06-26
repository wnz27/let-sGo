- [X] [网络轮询器](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-netpoller/)
  三种I/O模型，Go 对应epoll 源码解读
- [X] [Go语言基础之网络编程](https://www.liwenzhou.com/posts/Go/15_socket/)
    - 网络协议简介
    - 何为粘包
        > 主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。
        “粘包”可发生在发送端也可发生在接收端：
        由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
        接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。
    - 如何解决
        > 出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。
        封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。
        我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。
- [X] [HTTP 协议](https://hit-alibaba.github.io/interview/basic/network/HTTP.html)
- [X] [Improving web performance & security with TLS 1.3](https://www.cdn77.com/blog/improving-webperf-security-tls-1-3)
  看的太细致
- [X] [DNS 一般概览](https://cloud.google.com/dns/docs/dns-overview?hl=zh-cn)
  看的不细致
- [ ] [Android微信智能心跳方案](https://cloud.tencent.com/developer/article/1030660)
  懵懂
- [ ] [goim 架构与定制](https://juejin.cn/post/6844903827536117774)
  
- [ ] [如何设计一个亿级消息量的 IM 系统](https://xie.infoq.cn/article/19e95a78e2f5389588debfb1c)
- [ ] [Leaf：美团分布式ID生成服务开源](https://tech.meituan.com/2019/03/07/open-source-project-leaf.html)
- [ ] [系统调优，你所不知道的TIME_WAIT和CLOSE_WAIT](https://mp.weixin.qq.com/s/8WmASie_DjDDMQRdQi1FDg)
- [ ] [Java核心（五）深入理解BIO、NIO、AIO](https://www.imooc.com/article/265871)
- [ ] [从无到有：微信后台系统的演进之路](https://www.infoq.cn/article/the-road-of-the-growth-weixin-background)
- [ ] [DESIGN A CHAT SYSTEM](https://systeminterview.com/design-a-chat-system.php)
- [ ] [How Discord Stores Billions of Messages](https://blog.discord.com/how-discord-stores-billions-of-messages-7fa6ec7ee4c7)
- [ ] [](https://www.facebook.com/notes/facebook-engineering/the-underlying-technology-of-messages/454991608919/)
- [ ] [Flannel: An Application-Level Edge Cache to Make Slack Scale](https://slack.engineering/flannel-an-application-level-edge-cache-to-make-slack-scale/)
- [ ] [现代 IM 系统中的消息系统架构——模型篇](https://www.infoq.cn/article/emrual7ttkl8xtr-dve4)
- [ ] [如何打造千万级Feed流系统](http://www.91im.net/im/1130.html)

