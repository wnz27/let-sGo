/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/19 12:35 下午
* Description:
 */
package main

import "fmt"

const englishHelloPrefix = "Hello "

func Hello() string {
	return englishHelloPrefix + "world"
}

func HelloTo(name string) string {
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello())
}

