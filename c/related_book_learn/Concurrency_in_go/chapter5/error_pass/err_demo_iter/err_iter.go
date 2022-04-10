/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/29 20:52 8月
 **//**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/29 01:06 8月
 **/
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

// intermediate 模块

type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExcutable, err := isGloballyExec(jobBinPath)

	if err != nil {
		return IntermediateErr{wrapError(
			err,
			"cannot run job %q: requisite binaries not available",
			id,
		)} // 这里我们使用进行设计的异常信息。在这种情况下，我们想隐藏异常的底层细节，因为我们觉得这
		// 对我们模块的调用者来说并不重要。
	} else if isExcutable == false {
		return wrapError(
			nil,
			"cannot run job %q: requistite binaries are not executable",
			id,
		)
	}
	// 这里我们传递来自底层模块的异常。因为我们的体系结构决定，我们需要考虑从其他模块传递来的错误，
	// 而不是将它们用我们自己的错误类型封装，这里会存在一些问题，后面会提到。
	return exec.Command(jobBinPath, "--id=" + id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: \n", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	err := runJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."

		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(1 ,err, msg)
	}
}

