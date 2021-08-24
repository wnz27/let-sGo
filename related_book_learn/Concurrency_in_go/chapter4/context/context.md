# context 
正如我们所看到的，在并发程序中，由于超时，取消或系统其他部分的故障往往需要抢占操作。
我们已经看过了创建done channel 的习惯做法，该channel 在你的程序中流动并取消所有阻塞
的并发操作。这很好，但它也是有限的。

如果我们可以在简单的通知上附加传递额外的信息以取消：为什么会取消，或者我们的函数有一个
他必须要完成的最后期限，这将非常有用。
事实证明，对于任何规模的系统来说，使用这些信息来包装已完成的channel是非常常见的，因此
Go语言的作者们决定为此创建一个标准模式。
它起源于一个在标准库之外的实验功能，但是在Go1.7 中，context 包被引入标准库中，
这使它成为考虑并发问题时的一个标准的风格。

context 包：
```go
var Canceled = errors.New("context canceled")
var DeadlineExceeded error = deadlineExceededError{}

type CancelFunc func()
unc Background() Context
func TODO() Context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```
稍后讨论这些类型和函数，现在关注context 类型。
这个类型就像是 done channel 一样，在你的整个系统中进行传递。如果使用上下文包，
那么位于顶级并发调用下游的每个函数都会将context 作为其第一个参数。
类型如下：
```go
type Context interface {
	// Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
	Deadline() (deadline time.Time, ok bool)

	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	// The close of the Done channel may happen asynchronously,
	// after the cancel function returns.
	//
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	//
	// Done is provided for use in select statements:
	//
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
	// See https://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancellation.
	Done() <-chan struct{}

	// If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	// After Err returns a non-nil error, successive calls to Err return the same error.
	Err() error

	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	//
	// Use context values only for request-scoped data that transits
	// processes and API boundaries, not for passing optional parameters to
	// functions.
	//
	// A key identifies a specific value in a Context. Functions that wish
	// to store values in Context typically allocate a key in a global
	// variable then use that key as the argument to context.WithValue and
	// Context.Value. A key can be any type that supports equality;
	// packages should define keys as an unexported type to avoid
	// collisions.
	//
	// Packages that define a Context key should provide type-safe accessors
	// for the values stored using that key:
	//
	// 	// Package user defines a User type that's stored in Contexts.
	// 	package user
	//
	// 	import "context"
	//
	// 	// User is the type of value stored in the Contexts.
	// 	type User struct {...}
	//
	// 	// key is an unexported type for keys defined in this package.
	// 	// This prevents collisions with keys defined in other packages.
	// 	type key int
	//
	// 	// userKey is the key for user.User values in Contexts. It is
	// 	// unexported; clients use user.NewContext and user.FromContext
	// 	// instead of using this key directly.
	// 	var userKey key
	//
	// 	// NewContext returns a new Context that carries value u.
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 		return context.WithValue(ctx, userKey, u)
	// 	}
	//
	// 	// FromContext returns the User value stored in ctx, if any.
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	Value(key interface{}) interface{}
}
```
这看起来也简单，有一个Done 方法返回当我们的函数被强占时关闭的channel。
还有一些新的但易于理解的方法：一个Deadline 函数，用于指示在一定时间之后 goroutine 是否会被取消，
以及一个Err方法，如果goroutine被取消，将返回非零。但Value方法看起来似乎有点儿不合适。
这是为什么呢?

Go语言作者们注意到，goroutine 的主要用途之一是为请求提供服务的程序。
通常在这些程序中，除了强占信息之外，还需要传递特定于请求的信息。
这是Value函数的目的。我们稍后会进一步讨论这个问题，但现在我们只需要知道 context 包
有两个主要目的：
- 提供一个可以取消你的调用图中分支的API。
- 提供用于通过呼叫传输请求范围数据的数据包。

让我们先关注：取消。
正如我们在 "防止 goroutine 泄露" 中所学到的，函数中取消有三个方面：
- goroutine 的 父goroutine 可能想要取消它。
- 一个goroutine 可能想要取消它的子 goroutine。
- 【goroutine 中的任何阻塞操作都必须是可抢占的，以便它可以被取消。】

context包帮助管理所有这三个东西。

正如我们所提到的，context 类型将是你的函数的第一个参数。如果你看看 context 
接口上的方法，你会发现没有任何东西可以改变底层结构的状态。
此外，接收 context的函数并不能取消它。这保护了调用堆栈上的函数被子函数取消context的情况。
结合 done channel 提供的完成函数，这允许context 类型安全地管理其前件（没懂）的取消。

这就产生一个问题：如果 context 是不可变的，那么我们如何影响调用堆栈中当前函数下面的函数
中的取消行为？

下面的函数就是让 context 包变得如此重要的原因所在了。让我们从中再挑选出几个函数来刷新我们
的认知：
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
请注意，所有这些函数都接收一个Context 参数, 并且返回一个 Context。
其中一些还有其他的参数，如截止时间和超时参数。这些函数都使用与这些函数相关的
选项来生成Context的新实例。

WithCancel 返回一个新的Context，它在调用返回的cancel 函数时关闭其 done channel。
WithDeadline 返回一个新的Context，当机器的时钟超过给定的最后期限时，它关闭完成的channel。
WithTimeout 返回一个新的Context，它在给定的超时时间后关闭其完成的channel。

如果你的函数需要以某种方式在调用图中取消它后面的函数，它将调用其中一个函数并传递给它的上下文，
然后将返回的上下文传递给它的子元素。如果你的函数不需要修改取消行为，那么函数只传递给定的上下文。（这里没读懂翻译）

通过这种方式，调用图的连续图层可以创建符合其需求的上下文，而不会影响其父母节点。
这为如何管理调用图的分支提供了一个非常可组合的优雅解决方案。

context 包就是本着这种精神来串联起你程序的调用图的。在前面的对象的范例中，通常将对经常使用
的数据的引用存储为成员变量，但重要的是不要使用context.Context 的实例来执行此操作。context.Context 的实例
可能与外部看起来相同，但在内部它们可能会在每个栈帧更改。
出于这个原因，总是将context 的实例传递给你的函数是很重要的。通过这种方式，函数具有用于它的上下文，
而不是用于堆栈N 的context。

在异步调用图的顶部，你的代码可能不会传递context。要启动这个调用链，context包提供了两个函数来创建上下文的空实例。
```go
func Backgroud() Context
func TODO() Context
```
`Backgroud()`只是返回一个空的上下文。`TODO()` 不是用于生产，而是返回一个空的 context。`TODO()` 的预期目的
是作为一个占位符，当你不知道使用哪个 context，或者你希望你的代码被提供一个context，但上游代码还没有提供。

所以让我们把所有这些用于使用。我们来看一个使用完成channel模式的例子，并且看看我们切换到使用 context 包获得什么好处。
这是一个同时打印问候和告别的程序"
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Wait()
}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(3 * time.Second):
	}
	return "EN/US", nil
}
```
[【demo】](context_demo1/cd1_done.go)

输出：
```shell
goodbye world!
hello world!
```
忽略竞争条件（顺序不定）, 我们可以看到我们的程序有两个分支同时运行。
我们通过创建完成通道并将其传递给我们的调用图来设置标准抢占方法。
如果我们在main 的任何一点关闭done channel，那么两个分支都将被取消。

通过引入goroutine, 我们已经开辟了以几种不同且有趣的方式来控制该程序的可能性。
我们可能希望genGreeting 在耗时过长的时候发生超时。也许我们不希望genFarewell
调用locale，如果我们知道其父函数将很快被取消。
在每个堆栈框架中，一个函数可以影响其下的整个调用堆栈。

使用 done channel 模式的话，我们可以通过将输入的 done channel 都包装到统一的
done channel中，并在任意done channel 被关闭的时候关闭所有done channel，
【但是这样做的话我们没办法获得Context 所能给我们提供的关于错误以及超时的额外信息】。

为了比较 done channel 和 context 包，我们用树来表示这个程序，树中每个节点代表一个
函数的调用
![](call_1.jpg)

让我们修改我们的程序，使用 context 包而不是 done channel。因为我们现在
具有 context.Context 的灵活性，所以我们可以引入一个有趣的场景。

假设 genGreeting 只想在放弃调用 locale 之前 等待 1s，超时时间就应该为1s。
我们也想要在 main 函数中建立一些智能逻辑。如果 printGreeting 不成功，
我们也想取消我们对printFarewell 的调用。毕竟，如果我们不打声招呼，说再见就没有意义了。

使用context 包实现这一点很简单：
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 1、这里 main 用 context.Background() 创建一个 Context 并用 context.WithCancel 包装它以允许取消。
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			// 2、这一行上，如果从打印语问候语返回错误，main将取消这个context。
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()
	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {return err}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	// 3、这里genGreeting 用context.WithTimeout 包装它的Context。这将
	// 在1s 后自动取消返回的 context，从而取消它传递该 context 的任何子函数，即语言环境。
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		// 这一行返回为什么Context 被取消对的原因。该错误将会一直弹出到main， 这会导致取消。
		return "", ctx.Err()
	case <-time.After(30 * time.Second):
	}
	return "EN/US", nil
}
```
[【demo】](context_demo2/cd2_ctx.go)

输出：
```shell
cannot print greeting: context deadline exceeded
cannot print farewell: context canceled
```
让我们看使用调用图了解发生了什么。这里数字对应前面例子中的代码标注
![](call_2.jpg)

我们可以从我们的输出中看到该系统运行的很完美。
由于我们确保local 至 少需要一分钟来运行，因此我们在genGreeting 中的调用将
始终超时，这意味着main 会始终取消 printFarewell 下面的调用图。

请注意，genGreeting 是如何构建自定义的Context.Context 以满足其需求，
而不必影响父级的context。如果 genGreeting 成功返回，并且 printGreeting 
需要再次调用，则可以在不泄露有关genGreeting 如何操作的信息的情况下进行。
这种可组合性使你能够编写大型系统，而无需在整个调用图中混淆问题。

我们可以对这个程序进行另一个改进：因为我们知道locale 需要大约一分钟的时间来运行，
在locale的内部，我们可以检查是否已经设置了超时时间，如果已经设置了超时时间的话，
我们是否已经超时。下面例子展示了如何使用context.Context 的 Deadline 方法：
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {return err}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EV/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}


func locale(ctx context.Context) (string, error) {
	// 这里我们检查我们的context 是否提供了超时时间。如果确实如此，并且我们的系统时钟已经超过截止时间，
	// 那么我们只会返回context 包中定义的特定错误，即 DeadlineExceeded 
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1 * time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <- ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
```
输出
```shell
cannot print greeting: context deadline exceeded
cannot print farewell: context canceled
```
虽然在迭代程序之间的差异很小，但它允许locale 函数很快就会失败。
在调用下一个函数的成本很高的程序中，这可能会节省大量的时间，但
至少它也允许该函数立即失效，而不必等待实际的超时发生。
唯一的问题是，你必须知道你的下级调用图需要多长时间，这个实践起来可能非常困难。


