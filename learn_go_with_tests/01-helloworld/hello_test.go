/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/19 12:58 下午
* Description:
 */
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello world"

	if got != want {
		t.Errorf("got '%q' want '%q' ", got, want)
	}
}


func TestHelloTo(t *testing.T) {
	got := HelloTo("Alice")
	want := "Hello Alice"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}



