/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/23 6:46 下午
* Description:
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv.
	// This will return an empty string if the key isn’t present in the environment.
	os.Setenv("FOO", "1")
	// Running the program shows that we pick up the value for FOO that we set in the program,
	// but that BAR is empty.
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Use os.Environ to list all key/value pairs in the environment.
	// This returns a slice of strings in the form KEY=value.
	// You can strings.SplitN them to get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
	}

	/*
	$ go run environment-variables.go
	The list of keys in the environment will depend on your particular machine.

	If we set BAR in the environment first, the running program picks that value up.
	$ BAR=2 go run xxxxx.go
	output:
	FOO: 1
	BAR: 2

	....
	[BAR 2]
	[FOO 1]

	*/
}
