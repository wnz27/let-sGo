/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 5:59 下午
* Description:
 */
package main

/*
The sync package provides a safe mechanism for initialization in the presence of multiple goroutines
through the use of the Once type. Multiple threads can execute once.Do(f) for a particular f,
but only one will run f(), and the other calls block until f() has returned.

A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns.
 */

import (
	"sync"
	"time"
)

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main(){
	twoprint()
	time.Sleep(time.Millisecond * 500)
}
