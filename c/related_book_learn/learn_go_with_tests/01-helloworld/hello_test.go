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
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()  // ！！！！！会打出具体失败位置
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := HelloTo("Alice", "")
		want := "Hello Alice"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := HelloTo("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloTo("Elodie", "Spanish")
		want := "Hola Elodie"
		assertCorrectMessage(t, got, want)
	})
}



