/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/5 00:19 6月
 **/
package main

import "fmt"

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := app()
	b := app()
	a("go")

	fmt.Println(b("All"))

	fmt.Println(a("All"))

}
