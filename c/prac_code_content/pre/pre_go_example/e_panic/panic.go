/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/16 11:37 上午
* Description:
 */
package main

import "os"

func main() {
	//panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
