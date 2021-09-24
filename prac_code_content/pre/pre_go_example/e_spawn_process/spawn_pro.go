/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/24 11:59 上午
* Description:
 */
package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// We’ll start with a simple command that takes no arguments or input and just prints something to stdout.
	// The exec.Command helper creates an object to represent this external process.
	dateCmd := exec.Command("date")

	// .Output is another helper that handles the common case of running a command,
	// waiting for it to finish, and collecting its output.
	// If there were no errors, dateOut will hold bytes with the date info.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Next we’ll look at a slightly more involved case where we pipe data to the external process
	// on its stdin and collect the results from its stdout.
	grepCmd := exec.Command("grep", "hello")
	// Here we explicitly grab input/output pipes, start the process, write some input to it,
	// read the resulting output, and finally wait for the process to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	// 从grep输出收集结果
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// We omitted error checks in the above example,
	// but you could use the usual if err != nil pattern for all of them.
	// We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly delineated command and argument array,
	// vs. being able to just pass in one command-line string.
	// If you want to spawn a full command with a string, you can use bash’s -c option:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
