/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/19 12:35 下午
* Description:
 */
package main

import "fmt"

func Hello() string {
	return englishHelloPrefix + "world"
}

const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola "
const englishHelloPrefix = "Hello "
const frenchHelloPrefix = "Bonjour "

func HelloTo(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string)  (prefix string){
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello())
}

