/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/22 3:44 下午
* Description:
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// To start, here’s how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/tmp/dat2")
	check(err)
	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you’d expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.  刷盘，之前都在内存里。
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("worte %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()

	/*
	$ go run xxxx.go
	wrote 5 bytes
	wrote 7 bytes
	wrote 9 bytes

	Then check the contents of the written files.
	$ cat /tmp/dat1
	hello
	go

	$ cat /tmp/dat2
	some
	writes
	buffered
	 */
}
