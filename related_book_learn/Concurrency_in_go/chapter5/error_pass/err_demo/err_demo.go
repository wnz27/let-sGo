/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/29 01:06 8月
 **/
package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		// 我们存储了正在封装的异常。通常我们会希望能够找到最底层的异常，以便在需要时可以调查发生的异常。
		Inner: err,
		Message: fmt.Sprintf(messagef, msgArgs...),
		// 这行代码再创建异常时记录堆栈轨迹轨迹。过于复杂的错误类型经过wrapError 封装后可能会省略一些栈帧。
		StackTrace: string(debug.Stack()),
		// 我们创建一个可以存储各种杂志的变量。我们可以将并发ID，堆栈轨迹的hash 或可能有助于诊断异常的其他
		// 上下文信息存储在这里
		Misc: make(map[string]interface{}),
	}
}

// lowLevel 模块

type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		// 我们用自定义的异常来调用 os.Stat(path) 中的原始异常。
		// 在这种情况下，我们可以用这个异常传递信息，而不用对它做任何修饰。
		return false, LowLevelErr{wrapError(err, err.Error())}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

func (err MyError) Error() string {
	return err.Message
}

func main() {

}
