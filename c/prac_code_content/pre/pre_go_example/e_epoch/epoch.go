/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/18 12:28 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds,
	// respectively.
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	// Note that there is no UnixMillis,
	// so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//  You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))


}
