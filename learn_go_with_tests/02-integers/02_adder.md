# 整数 

## 首先编写测试
```go
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
```
你会发现，我们使用 %d 而不是 %s 作为格式字符串。这是因为我们希望打印一个整数而不是字符串。

还需要注意的是，我们不再使用 main 包，而是定义了一个名为 integers 的包。
顾名思义，它将聚集那些处理整数的函数，例如 Add。
执行 go test 运行测试
检查编译错误: `./adder_test.go:6:9: undefined: Add`

## 为测试的运行编写最少量的代码并检查失败测试的输出
编写足够的代码来使编译通过，仅此而已 —— 请记住，我们要查看的是测试是否因为合理的原因失败。
```go
package integers

func Add(x, y int) int {
	return 0
}
```
当你有多个相同类型的参数（在我们的例子中是两个整数），可以将它缩短为 (x，y int) 而不是 (x int, y int)。
现在运行测试，我们应该很高兴测试正确报告了错误信息。
```go
--- FAIL: TestAdder (0.00s)
adder_test.go:16: expected 4 but got 0
```

如果你注意一下就会发现，我们在 上一节 中学习了 命名返回值，但没有在这里使用。
**它通常应该在结果的含义在上下文中不明确时使用**，在我们的例子中，Add 函数将参数相加是非常明显的。





