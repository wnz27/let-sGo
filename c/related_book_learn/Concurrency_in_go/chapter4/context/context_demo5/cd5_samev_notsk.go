/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/26 01:23 8月
 **/
package main

import "fmt"

type foo int
type bar int

func main() {
	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(1)] = 2
	fmt.Printf("%v\n", m)
}
