## reference: [一起聊聊 Go Context 的正确使用姿势](https://mp.weixin.qq.com/s?__biz=Mzg3NTU3OTgxOA==&mid=2247492385&idx=1&sn=2f7d5cf6a7d34706b6decb1fd18b0afe&chksm=cf3df3e4f84a7af2dcfd040095165c862a5cc36d516cfca5f8c3e8caa21b8795044987af8a0c&mpshare=1&scene=1&srcid=0720QK70CQOD2J5QkHz2BvOa&sharer_sharetime=1626768045982&sharer_shareid=d94ad27d4946e2a1fa2bda2006d8985f&version=3.1.10.90255&platform=mac#rd)

## 正确的使用姿势
- 对第三方调用要传入 context
- 不要将上下文存储在结构类型中
> 当然，我们也不能一杆子打死所有情况。确实存在极少数是把 context 放在结构体中的。基本常见于：底层基础库。 DDD 结构。
- 函数调用链必须传播上下文
> 我们会把 context 作为方法首位，本质目的是为了传播 context，自行完整调用链路上的各类控制

## context 的继承和派生
在 Go 标准库 context 中具有以下派生 context 的标准方法
- func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
- func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
- func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

一般会有父级 context 和子级 context 的区别，我们要保证在程序的行为中上下文对于多个 goroutine 同时使用是安全的。
并且存在父子级别关系，父级 context 关闭或超时，可以继而影响到子级 context 的程序。

## 函数调用链必须传播上下文
我们会把 context 作为方法首位，本质目的是为了传播 context，自行完整调用链路上的各类控制：
```go
func List(ctx context.Context, db *sqlx.DB) ([]User, error) {
 ctx, span := trace.StartSpan(ctx, "internal.user.List")
 defer span.End()

 users := []User{}
 const q = `SELECT * FROM users`

 if err := db.SelectContext(ctx, &users, q); err != nil {
  return nil, errors.Wrap(err, "selecting users")
 }

 return users, nil
}
```
像在上述例子中，我们会把所传入方法的 context 一层层的传进去下一级方法。
这里就是将外部的 context 传入 List 方法，再传入 SQL 执行的方法，解决了 SQL 执行语句的时间问题。

## 不传递 nil context
很多时候我们在创建 context 时，还不知道其具体的作用和下一步用途是什么。
这种时候大家可能会直接使用 context.Background 方法：
```go
var (
   background = new(emptyCtx)
   todo       = new(emptyCtx)
)

func Background() Context {
   return background
}

func TODO() Context {
   return todo
}
```
但在实际的 context 建议中，我们会建议使用 context.TODO 方法来创建顶级的 context，
直到弄清楚实际 Context 的下一步用途，再进行变更。

## context 仅传递必要的值

我们在使用 context 作为上下文时，经常有信息传递的诉求。
像是在 gRPC 中就会有 metadata 的概念，而在 gin 中就会自己封装 context 作为参数管理。
Go 标准库 context 也有提供相关的方法：
```go
type Context
func WithValue(parent Context, key, val interface{}) Context
```
使用如下
```go
func main() {
 type favContextKey string
 f := func(ctx context.Context, k favContextKey) {
  if v := ctx.Value(k); v != nil {
   fmt.Println("found value:", v)
   return
  }
  fmt.Println("key not found:", k)
 }

 k := favContextKey("脑子进")
 ctx := context.WithValue(context.Background(), k, "煎鱼")

 f(ctx, k)
 f(ctx, favContextKey("小咸鱼"))
}
```
输出：
```shell
found value: 煎鱼
key not found: 小咸鱼
```

在规范中，我们建议 context 在传递时，仅携带必要的参数给予其他的方法，
或是 goroutine。甚至在 gRPC 中会做严格的出、入上下文参数的控制。
在业务场景上，context 传值适用于传必要的业务核心属性，
例如：租户号、小程序ID 等。
不要将可选参数放到 context 中，否则可能会一团糟。

##总结
- 对第三方调用要传入 context，用于控制远程调用。
- 不要将上下文存储在结构类型中，尽可能的作为函数第一位形参传入。
- 函数调用链必须传播上下文，实现完整链路上的控制。
- context 的继承和派生，保证父、子级 context 的联动。
- 不传递 nil context，不确定的 context 应当使用 TODO。
- context 仅传递必要的值，不要让可选参数揉在一起。
