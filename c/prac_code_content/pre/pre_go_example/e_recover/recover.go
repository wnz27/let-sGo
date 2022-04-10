/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/16 12:00 下午
* Description:
 */
package main

import "fmt"

/*
An example of where this can be useful: a server wouldn’t want to crash if one of the client connections
exhibits a critical error.
Instead, the server would want to close that connection and continue serving other clients.
In fact, this is what Go’s net/http does by default for HTTP servers.
 */

func myPanic() {
	panic("my panic!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: \n", r)
		}
	}()

	myPanic()

	fmt.Println("After my panic()")
}
