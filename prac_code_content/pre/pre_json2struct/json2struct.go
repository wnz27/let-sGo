/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/11 4:16 下午
* Description:
 */
package main

type social struct {
	facebook string
	twitter string
}

type User struct {
	Name string
	Type string
	Age int
	Social social
}
const jsonFile = "./pre/pre_json2struct/user.json"
func main() {
}
