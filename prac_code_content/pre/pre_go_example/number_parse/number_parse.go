/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/21 18:50 9月
 **/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.2345", 64)
	fmt.Println(f)

	// For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// A ParseUint is also available.
	u, _ := strconv.ParseUint("789", 0 , 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("146")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("qer")
	fmt.Println(e)
}
