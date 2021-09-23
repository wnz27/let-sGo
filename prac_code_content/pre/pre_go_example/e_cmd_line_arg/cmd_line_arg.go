/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/22 11:26下午
* Description:
 */
package main

import (
	"fmt"
	"os"
)

/*
Command-line arguments are a common way to parameterize execution of programs.
For example, go run hello.go uses run and hello.go arguments to the go program.
 */


func main() {
	// os.Args provides access to raw command-line arguments.
	// Note that the first value in this slice is the path to the program,
	// and os.Args[1:] holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//You can get individual args with normal indexing.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	/*

	To experiment with command-line arguments it’s best to build a binary with go build first.

	$ go build prac_code_content/pre/pre_go_example/e_cmd_line_arg/cmd_line_arg.go
	$ ./cmd_line_arg a b c d

	output:
	[./cmd_line_arg a b c d]
	[a b c d]
	c

	*/
}
