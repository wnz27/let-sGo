/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/19 22:15 4月
 **/
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// 自定义的出错结构
type myError struct {
	arg int
	errMsg string
}

// 实现Error 方法
func (e *myError) Error() string {
	return fmt.Sprintf("%d inner-split %s", e.arg, e.errMsg)
}

// 两种出错
func error_test(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Bad Arg - negative!")
	} else if arg > 256 {
		return -1, &myError{arg: arg, errMsg: "Bad Arg - too large!"}
	}
	return arg * arg, nil
}


func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func fn1() error {
	return errors.Errorf("123")
}

/*
When you have an error message that requires formatting, use the Errorf function from the fmt package:
var ErrInvalidParam = fmt.Errorf("invalid parameter [%s]", param)
 */

func main() {
	err := fn()
	//fmt.Printf("%+v", err)
	// 输出：这样可以把整个栈的信息留住
	/*
	error
	main.fn
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:27
	main.main
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:34
	runtime.main
	        /usr/local/go/src/runtime/proc.go:225
	runtime.goexit
	        /usr/local/go/src/runtime/asm_amd64.s:1371
	inner
	main.fn
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:28
	main.main
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:34
	runtime.main
	        /usr/local/go/src/runtime/proc.go:225
	runtime.goexit
	        /usr/local/go/src/runtime/asm_amd64.s:1371
	middle
	main.fn
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:29
	main.main
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:34
	runtime.main
	        /usr/local/go/src/runtime/proc.go:225
	runtime.goexit
	        /usr/local/go/src/runtime/asm_amd64.s:1371
	outer
	main.fn
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:30
	main.main
	        /Users/fzk27/fzk27/let-sGo/code_content/go_error/go_error.go:34
	runtime.main
	        /usr/local/go/src/runtime/proc.go:225
	runtime.goexit
	        /usr/local/go/src/runtime/asm_amd64.s:1371
	*/
	fmt.Printf("%s, %v", err, err)  // output: 	outer: middle: inner: error, outer: middle: inner: error

	for _, i := range []int{-1, 4, 1000} {
		if r, e := error_test(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("success:", r)
		}
	}

	// output:
	// failed: Bad Arg - negative!
	// success: 16
	// failed: 1000 inner-split Bad Arg - too large!

	aaa := fn1()
	fmt.Println("\n================================")
	fmt.Println(errors.Cause(aaa))
	fmt.Println("\n================================")
	fmt.Printf("%+v", aaa)
	fmt.Println("\n================================")
	fmt.Println(errors.Cause(err))
}
