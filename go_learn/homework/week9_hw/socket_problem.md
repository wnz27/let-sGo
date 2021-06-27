## 总结几种 socket 粘包的解包方式, 尝试举例应用
- fix length 固定缓冲区大小
  - 对消息内容长度有足够掌握的情况
- delimiter based 界限分隔符
  - 天然适用于按行分割的消息, 如果我们基本可以确定消息最大长度
- length field based frame decoder  封装协议
  - 适用于消息是不定长的场景

粘包情形
- [server_code](./origin_server_client/server_hw/hw_s_origin.go)
- [client_code](./origin_server_client/client_hw/hw_c_origin.go)


## 实现一个从 socket connection 中解码出 goim 协议的解码器。
- [goim 协议长这样](http://goim.io/docs/protocol.html)
- [编解码器](./goim_decoder_attempt/goim_decoder.go)
- 实际尝试
  - [client](./goim_decoder_attempt/mock_goim_protocal_c/mock_goim_p_c.go)
  - [server](./goim_decoder_attempt/decode_goim_s/decode_goim_s.go)


