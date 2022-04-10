# DI 容器
golang 现有的依赖注入框架:
- [使用反射实现的-uber-dig](https://github.com/uber-go/dig)
- [使用 generate 实现的-google-wire](https://github.com/google/wire)

这里将通过反射实现一个类似 dig 简单的 demo，还可以通过读取配置文件，然后进行生成，
下面提供是通过 provider 进行构建依赖关系。

[【demo】](di_demo.go)
我们这里的实现比较粗糙，但是作为一个 demo 理解 di 容器也足够了，和 dig 相比还缺少很多东西，并且有许多的问题，例如 依赖关系，一种类型如果有多个 provider 如何处理等等等等。
可以看到我们总共就三个函数
- Provide: 获取对象工厂，并且使用一个 map 将对象工厂保存
- Invoke: 执行入口
- buildParam: 核心逻辑，构建参数
    - 从容器中获取 provider
    - 递归获取 provider 的参数值
    - 获取到参数之后执行函数
    - 将结果缓存并且返回结果






