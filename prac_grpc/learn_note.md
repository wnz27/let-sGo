RPC —— 调用规约
• 参数是一个结构体指针
• 返回值是一个结构体指针和error
• 参数放在 HTTP body 里面进行传输，采 用JSON作为序列化格式

RPC —— 如何解决 endpoint 写死的问题?
• 引入配置模块
• 引入服务名的机制，用服务名作为服务的 唯一性ID(名称)
• 不同服务会有不同的配置
• 维持一个 服务名 => 服务配置的映射

反射：
```
func SetFuncField2(val interface{}) {
	v := reflect.ValueOf(val) // 这是指针的反射
	ele := v.Elem()           // 拿到了指针指向的结构体
	t := ele.Type()           // 拿到了指针指向的结构体的类型信息
```

## Golang 语法——类型断言
• 形式:t, ok := x.(T) 或者 t := x.(T)
• T 可以是结构体或者指针
• 如何理解?
• 即x是不是T。
• 类似Java instanceOf + 强制类型转
换合体
• 如果 x 是 nil，那么永远是 false • 编译器不会帮你检查

## Golang 语法——类型转换 
• 形式:y := T(x)
• 如何理解?记住数字类型转换，string 和 []byte 互相转
• 类似Java强制类型转换
• 编译器会进行类型检查，不能转换的会编 译错误


## Golang 语法 —— map
• 基本形式:map[KeyType]ValueType • 创建 make 命令，或者直接初始化
• 取值:val, ok := m[key]
• 设值:m[key]=val
• key 类型:“可比较”类型



