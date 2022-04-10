<!--
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-03-21 14:58:44
 * @LastEditTime: 2022-04-10 17:36:05
 * @FilePath: /let-sGo/source_code_read/go-ctx-1.16.10/context-learn.md
 * @description: type some description
-->
# context 包 1.16.10

```go
// Package context defines the Context type, which carries deadlines,
// cancellation signals, and other request-scoped values across API boundaries and between processes.
```
定义了 Context 上下文 类型，包含了截止时间，取消信号，其他跨 API 界限的请求范围值，以及进程之间。

```go
// Incoming requests to a server should create a Context, and outgoing
// calls to servers should accept a Context. The chain of function
// calls between them must propagate the Context, optionally replacing
// it with a derived Context created using WithCancel, WithDeadline,
// WithTimeout, or WithValue. When a Context is canceled, all
// Contexts derived from it are also canceled.
```
传入服务器的请求应该创建一个Ctx，对服务器的传出调用应该接收一个上下文（个人理解是要设计一个ctx 参数）。
他们之间的函数调用链必须传播context，但是可以用Context的衍生来替代，比如使用 `WithCancel`, `WithDeadline`,
`WithTimeout`, 或者`WithValue`。


```go
// The WithCancel, WithDeadline, and WithTimeout functions take a
// Context (the parent) and return a derived Context (the child) and a
// CancelFunc. Calling the CancelFunc cancels the child and its
// children, removes the parent's reference to the child, and stops
// any associated timers. Failing to call the CancelFunc leaks the
// child and its children until the parent is canceled or the timer
// fires. The go vet tool checks that CancelFuncs are used on all
// control-flow paths.
```
这几个方法 `WithCancel`, `WithDeadline`, 和 `WithTimeout` 需要父Context 为参数，
返回一个衍生的Context 和一个取消函数。调用取消函数取消子ctx以及他们的子ctx，
移除父ctx 对子ctx 的引用，还有停下所有关联的计时器。
调用取消函数失败会泄露子及他们的子ctx，直到父ctx 被取消或者计时器触发。
go vet 工具检查是否在所有控制流路径上使用了取消函数。

```go
// Programs that use Contexts should follow these rules to keep interfaces
// consistent across packages and enable static analysis tools to check context
// propagation:
```
使用Context 应该遵循以下的规则来保证跨包的接口一致，并允许静态分析工具来检查 context 的传播

```go
// Do not store Contexts inside a struct type; instead, pass a Context
// explicitly to each function that needs it. The Context should be the first
// parameter, typically named ctx:
//
// 	func DoSomething(ctx context.Context, arg Arg) error {
// 		// ... use ctx ...
// 	}
```
1、不要在 struct 中存 Context；相反，如果需要的话，明确的在每个函数把 Context 当做参数来传播。
Context 应该放在第一个参数，通常名字为 ctx。

```go
// Do not pass a nil Context, even if a function permits it. Pass context.TODO
// if you are unsure about which Context to use.
```
2、不要传 nil Context， 即使这个方法允许. 当你不确定使用哪种 Context 时，使用 `context.TODO`
传递。

```go
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
```
3、仅在请求api范围中或者传输过程中使用Context 传值，不要用Context 传递可选参数

```go
// The same Context may be passed to functions running in different goroutines;
// Contexts are safe for simultaneous use by multiple goroutines.
```
4、同一个Context 也许会传给不同运行的 `goroutines` 的函数，Context 可以安全的同时被多个 groutines 使用


```go
// See https://blog.golang.org/context for example code for a server that uses
// Contexts.
```
浏览 https://blog.golang.org/context 示例代码看服务如何使用 Context

