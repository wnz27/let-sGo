/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/21 8:18 下午
* Description:
 */
package main

import "fmt"

func main()  {
	a := []int{1, 2, 3}
	b := a[:0]
	fmt.Printf("%o, %p, %p", b, &b, &a)
}
