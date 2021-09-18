/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/16 4:03 下午
* Description:
 */
package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:   ", s.Contains("test", "es"))
	p("Count:      ", s.Count("test", "t"))
	p("HasPrefix:   ", s.HasPrefix("test", "te"))
	p("HasSuffix:   ", s.HasSuffix("test", "st"))
	p("Index:      ", s.Index("test", "e"))
	p("Join:       ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:     ", s.Repeat("a", 5))
	p("Replace:    ", s.Replace("foo", "o", "0", 1))  // 替换1个
	p("Replace:    ", s.Replace("foo", "o", "0", -1))  // 全部
	p("Replace:    ", s.Replace("foo", "o", "0", 3))  // 超过数量没事
	p("Split:      ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:    ", s.ToLower("TEST"))
	p("ToUpper:    ", s.ToUpper("test"))
	p()

	p("Len:        ", len("hello"))
	p("Char:       ", "hello"[1])
}
