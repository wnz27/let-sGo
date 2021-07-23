# RPC
remote process call

人话，从客户端把service名，方法名，参数弄（序列化，编码）成一坨二进制流，通过网络，到
服务端，客户端翻译（反序列化，解码）出来, 然后拼成需要的对象方法，执行完结果再
弄一坨二进制，返回客户端，客户端解码拿到结果。

## gRpc

人话，google开发的，基于HTTP/2.0 协议之上的跨网络调用。

### 四种服务类型
1、Unary, 客户端发请求, server回应
2、Client-side streaming 用户以流的形式不断上传请求到server，
结束之后，server只返回一个响应
3、Server-side streaming 用户只上传一个请求，server源源不断的把
流传回给客户端
4、Bidirectional streaming 双工流形式交互

### TODO
描述一个向导服务，服务名称为RouteGuide:
- [X] 定义四种不同的信息类型分别为Point，Rectangle, Feature, RouteSummary以及Chat
定义四个方法：
- [X] GetFeature(输入一个Point) 返回这个点的Feature
- [X] ListFeatures(输入一个Rectangle) 输出流这个区域所有的Feature)
- [X] RecordRoute(输入流为每个时间点的位置Point) 返回一个RouteSummary
- [X] Recommend(输入流RecommendationRequest) 输出流Feature

- [proto文件](grpc_demo/route/route.proto)
- [桩代码生成](grpc_demo/route/gen.sh)
- [client - 暂无]()
- [server - 暂无]()



