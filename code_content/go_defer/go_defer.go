/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/20 02:06 4月
 **/
package main

import (
	"io"
	"os"
	"fmt"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func hello(i int) {
	fmt.Println(i)
}

func main() {

	i := 5
	defer hello(i)
	i = i + 10
	/*
	这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.
	 */


}

