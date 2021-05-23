# 数组与切片

数组允许你以特定的顺序在变量中存储相同类型的多个元素。
对于数组来说，最常见的就是迭代数组中的元素。
我们创建一个 Sum 函数，它使用 for 来循环获取数组中的元素并返回所有元素的总和。
让我们使用 TDD 思想。

## 先写测试函数
```go
package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15
	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
```

数组的容量是我们在声明它时指定的固定值。我们可以通过两种方式初始化数组：
```go
[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}
```

在错误信息中打印函数的输入有时很有用。我们使用 %v（默认输出格式）占位符来打印输入，它非常适用于展示数组。

## 运行测试
```go
FAIL    fzkprac/learn_go_with_tests/04-arrays [build failed]
```
## 先使用最少的代码来让失败的测试先跑起来：
`sum.go`
```go
package main

func Sum(numbers [5]int) (sum int) {
	return 0
}

```
测试还是会失败
```go
--- FAIL: TestSum (0.00s)
    sum_test.go:17: got 0 want 15 given, [1 2 3 4 5]
FAIL
exit status 1
FAIL    fzkprac/learn_go_with_tests/04-arrays   0.714s
```
## 把代码补充完整，使得它能够通过测试：
```go
package main

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return
}
```
可以使用 array[index] 语法来获取数组中指定索引对应的值。
在本例中我们使用 for 循环分 5 次取出数组中的元素并与 sum 变量求和。

## 一个源码版本控制的小贴士
此时如果你正在使用源码的版本控制工具（你应该使用它！），我会在此刻先提交一次代码。因为我们已经拥有了一个有测试支持的程序。

但我 不会 将它推送到远程的 master 分支，因为我马上就会重构它。
在此时提交一次代码是一种很好的习惯。因为你可以在之后重构导致的代码乱掉时回退到当前版本。

你总是能够回到这个可用的版本。

