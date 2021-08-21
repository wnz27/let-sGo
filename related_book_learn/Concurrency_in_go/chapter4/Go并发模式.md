# 第四章学习内容
本章深入将上一章的基本元组合成模式，以帮助保持系统的可扩展性和可维护性。

但是开始之前做一些解释，谈谈本章所包含的一些模式的格式。
在很多的示例中，我们将使用传递空接口(interface{}) 的channel。
在Go语言中使用空接口是有争议的，不过我们出于以下原因选择使用空接口。
- 首先，它使得在书中的其余部分编写简洁的例子变得更容易。
- 其次，在某些情况下，我认为这更能代表该模式正在努力达成的目标。 我们将在本章后面的"pipeline"直接讨论这一点。

如果这对你来说过于难以接受，请记住你始终可以为此代码创建Go语言生成器，并生成模式以利用你感兴趣的类型。

- [约束](constrain/约束.md)
- [for-select](for_select.md)
- [防止goroutine泄露](goroutine_leak/goroutine_leak.md)
- [or-channel](or_channel/or_channel.md)
    - [ ] TODO 可以尝试把 or 函数时间复杂度改为 logn ？
- [错误处理](err_handle/err_handle.md)
- [pipeline](pipeline/pipeline.md)
- [构建 pipeline 的最佳实践](pipeline_prac/pipeline_prac.md)
- [一些便利的生成器](generator/generator.md)
- [扇出，扇入](fan_out_fan_in/fan_in_out.md)
- [or-done-channel](or_done_channel/or-done-channel.md)
- [tee-channel](tee_channel/tee_channel.md)
- [桥接 channel](bridge_channel/bridge_channel.md)
- [队列排队](queue_c/q_c.md)

