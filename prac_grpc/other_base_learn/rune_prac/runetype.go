/**
 * @project prac_grpc
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/6 11:16 6月
 **/
package main

import (
	"fmt"
	"unicode/utf8"
)

func runeType() {
	a := "你好啊"
	fmt.Printf("a len: %d", utf8.RuneCountInString(a))
}

func main() {
	runeType()
}