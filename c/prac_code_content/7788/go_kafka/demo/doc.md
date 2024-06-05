<!--
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-06-04 14:19:43
 * @LastEditTime: 2024-06-05 11:43:43
 * @FilePath: /let-sGo/c/prac_code_content/7788/go_kafka/demo/doc.md
 * @description: type some description
-->
# 部署 demo 应用
## 创建 ns
```shell
k create ns event-demo
```
创建事件示例名称空间，然后将knative-eventing-injection标签添加到该名称空间。您可以使用名称空间将它们组合在一起并组织您的Knative资源，包括Eventing子组件
```shell
k label namespace event-demo knative-eventing-injection=enabled
```

## web server
```shell
k apply -f web.yaml -n env-feature-docking-gmold
k apply -f web.yaml -n op-demo
```

## 启动一个可以发 curl 的 pod
```shell
k apply -f curl_pod.yaml -n event-demo
```

## 确认 broker url（已有的 broker)
```shell
k describe Broker -n target_ns
```
找到 Address:URL
`http://kafka-broker-ingress.knative-eventing.svc.cluster.local/env-feature-docking-gmold/kafka-broker`

## 建立一个 trigger 使 trigger 可以转发相应规则的 event 到 web server
```yaml
apiVersion: eventing.knative.dev/v1
kind: Trigger
metadata:
  name: h-trigger
spec:
    broker: kafka-broker
    filter:
      attributes:
        type: greeting
        source: send
    subscriber:
      ref: 
        apiVersion: v1
        kind: Service
        name: hello-display
```
创建
```shell
k apply -f trigger_demo.yaml -n env-feature-docking-gmold
```

## 往 broker 发送消息
进入 curl 的 pod
往 broker 发消息
```shell
curl -v "http://kafka-broker-ingress.knative-eventing.svc.cluster.local/env-feature-docking-gmold/kafka-broker" \
  -X POST \
  -H "Ce-Id: say-hello" \
  -H "Ce-Specversion: 0.3" \
  -H "Ce-Type: greeting" \
  -H "Ce-Source: not-sendoff" \
  -H "Content-Type: application/json" \
  -d '{"msg":"12348123740912347jsdfaskdjfhasjklfhdslkfashdlfkshfkl"}'


curl -v "http://kafka-broker-ingress.knative-eventing.svc.cluster.local/env-feature-docking-gmold/kafka-broker" \
  -X POST \
  -H "Ce-Id: say-hello" \
  -H "Ce-Specversion: 0.3" \
  -H "Ce-Type: greeting" \
  -H "Ce-Source: send" \
  -H "Content-Type: application/json" \
  -d '{"msg":"Hello Kafka!!!!"}'

  curl -v "localhost:8080/msg/receive" \
  -X POST \
  -H "Ce-Id: say-hello" \
  -H "Ce-Specversion: 0.3" \
  -H "Ce-Type: greeting2" \
  -H "Ce-Source: send2" \
  -H "Content-Type: application/json" \
  -d '{"msg":"Hello other service !!!!"}'

  curl -v "https://hello-event.newtest.k8s.guanmai.cn/msg/receive" \
  -X POST \
  -H "Ce-Id: say-hello" \
  -H "Ce-Specversion: 0.3" \
  -H "Ce-Type: greeting2" \
  -H "Ce-Source: send2" \
  -H "Content-Type: application/json" \
  -d '{"msg":"Hello other service !!!!"}'


  curl -v "http://kafka-broker-ingress.knative-eventing.svc.cluster.local/env-feature-docking-gmold/kafka-broker" \
  -X POST \
  -H "Ce-Id: say-hello" \
  -H "Ce-Specversion: 0.3" \
  -H "Ce-Type: greeting2" \
  -H "Ce-Source: send2" \
  -H "Content-Type: application/json" \
  -d '{"msg":"Hello other service !!!!"}'
```

## 测试 trigger 使用 uri 配置
在测试集群部署一个 web 服务
启动第二个 trigger
```shell
k apply -f trigger_demo2.yaml -n env-feature-docking-gmold
```

### curl 服务
```shell
curl -v "http://hello-event.newtest.k8s.guanmai.cn/" \
  -X POST \
  -H "Ce-Id: say-hello" \
  -H "Ce-Specversion: 0.3" \
  -H "Ce-Type: greeting" \
  -H "Ce-Source: send" \
  -H "Content-Type: application/json" \
  -d '{"msg":"Hello Kafka!!!!"}'
```

## 部署 trigger + sink
```shell
k apply -f wms_demo.yaml -n env-feature-docking-gmold
```


