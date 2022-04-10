/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/16 11:42 上午
* Description:
 */
package main

import (
	"fmt"
	"os"
)

func createFile(fileAbsPath string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(fileAbsPath)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
