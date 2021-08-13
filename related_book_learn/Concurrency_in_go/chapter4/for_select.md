# for-select 循环
在go语言中你会一遍又一遍地看到for-select循环。它不过是这样的：
```go
for {  // 要不就无限循环，要不就使用 range 语句循环
    select {
    // 使用channel进行作业
    }
}
```
有以下几种情况你可以见到这种模式。

### 向channel发送迭代变量
通常情况下，你需要将可迭代的内容转换为channel上的值。这不是什么幻想，通常看起来像这样：
```go
for _, s := range []string{"a", "b", "c"} {
    select {
    case <- done:
        return
    case stringStream <- s:
    }
}
```
### 循环等待停止
创建循环，无限循环直到停止的goroutine很常见。这个有一些变化。
选择哪一种纯属是一种个人爱好。

第一种变体保持select 语句尽可能短：
```go
for{
	select {
	case <- done:
		return
    default:
    }
    // 进行非抢占式任务
}
```
如果已经完成的channel未关闭，我们将退出select语句并继续执行for循环
