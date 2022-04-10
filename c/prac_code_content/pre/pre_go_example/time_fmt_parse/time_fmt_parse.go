/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/18 2:24 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	// Here’s a basic example of formatting a time according to RFC3339, using the corresponding layout constant.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Time parsing uses the same layout values as Format.
	t1, e := time.Parse(
		time.RFC3339,
		"2021-09-18T15:07:41+00:00")
	p(t1, "Err:", e)

	p(t.Format("3:04PM"))
	p(t.Format("Sat Sep _2 15:07:05 2021"))
	p(t.Format("2021-09-18T15:07:05.999999-07:00"))
	form := "3 07 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// For purely numeric representations you can also use standard string formatting with
	// the extracted components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse will return an error on malformed input explaining the parsing problem.
	ansic := "Sat Sep 18 15:07:06 2021"
	_, e = time.Parse(ansic, "3:07PM")
	p(e)
}
