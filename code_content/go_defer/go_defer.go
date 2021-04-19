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

func main() {

}

