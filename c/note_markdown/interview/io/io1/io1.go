/**
 * @project let-sGo
 * @Author 27
 * @Description //
 * @Date 2021/7/25 15:52 7月
 **/
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	buffer := make([]byte, 1024)
	readerFromBytes := bytes.NewReader(buffer)
	n, err := io.Copy(ioutil.Discard, readerFromBytes)
	// n == 1024, err == nil
	fmt.Printf("n=%v,err=%v\n", n, err)

	data := "hello world"
	readerFromBytes2 := strings.NewReader(data)
	n2, err2 := io.Copy(ioutil.Discard, readerFromBytes2)
	fmt.Printf("n=%v,err=%v\n",n2, err2)

	// 一行代码实现回显
	io.Copy(os.Stdout, os.Stdin)
}
