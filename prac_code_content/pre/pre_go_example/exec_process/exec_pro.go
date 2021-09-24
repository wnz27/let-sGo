/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/24 12:00 下午
* Description:
 */
package main

import (
	"os"
	"os/exec"
	"syscall"
)

/*

In the previous example we looked at spawning external processes.
We do this when we need an external process accessible to a running Go process.

Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one.
To do this we’ll use Go’s implementation of the classic exec function.
*/

func main() {
	// For our example we’ll exec ls. Go requires an absolute path to the binary we want to execute,
	// so we’ll use exec.LookPath to find it (probably /bin/ls).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form (as opposed to one big string).
	// We’ll give ls a few common arguments. Note that the first argument should be the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use. Here we just provide our current environment.
	env := os.Environ()

	// Here’s the actual syscall.Exec call. If this call is successful,
	// the execution of our process will end here and be replaced by the /bin/ls -a -l -h process.
	// If there is an error we’ll get a return value
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	/*
	When we run our program it is replaced by ls.

	Note that Go does not offer a classic Unix fork function. Usually this isn’t an issue though,
	since starting goroutines, spawning processes, and exec’ing processes covers most use cases for fork.
	 */
}
