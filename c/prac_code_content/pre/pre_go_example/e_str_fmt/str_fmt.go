/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/16 5:25 下午
* Description:
 */
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {
	p := point{1, 2}
	// 打出结构体 不带变量名只有变量值
	pf("struct1: %v\n", p)

	// 打印出结构体，带变量名和变量值
	// If the value is a struct, the %+v variant will include the struct’s field names.
	pf("struct2: %+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e.
	// the source code snippet that would produce that value.
	pf("struct3: %#v\n", p)

	// To print the type of a value, use %T.
	pf("type: %T\n", p)

	// Formatting booleans is straight-forward.
	pf("bool: %t\n", true)

	// There are many options for formatting integers. Use %d for standard, base-10 formatting.
	pf("int: %d\n", 123)

	// This prints a binary representation.
	pf("bin: %b\n", 14)

	// This prints the character corresponding to the given integer.
	pf("char: %c\n", 33)

	// %x provides hex encoding.
	pf("hex: %x\n", 456)

	// There are also several formatting options for floats.
	// For basic decimal formatting use %f.
	pf("float1: %f\n", 78.9)

	// %e and %E format the float in (slightly different versions of) scientific notation.
	pf("float2: %e\n", 123400000.0)
	pf("float3: %E\n", 123400000.0)

	// For basic string printing use %s.
	pf("str1: %s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.
	pf("str2: %q\n", "\"string\"")

	// As with integers seen earlier,
	// %x renders the string in base-16, with two output characters per byte of input.
	pf("str3: %x\n", "hex this")

	// To print a representation of a pointer, use %p.
	pf("pointer: %p\n", &p)

	// When formatting numbers you will often want to control the width and precision of the resulting figure.
	// To specify the width of an integer, use a number after the % in the verb.
	// By default the result will be right-justified and padded with spaces
	pf("width1: |%6d|%6d|\n", 12222, 345)

	// You can also specify the width of printed floats, though usually you’ll also want to
	// restrict the decimal precision at the same time with the width.precision syntax.
	pf("width2: |%6.2f|%6.2f|\n", 12.2252, 3.454)

	// To left-justify, use the - flag.
	pf("width3: |%-6.2f|%-6.2f|\n", 13.2345, 3.456)

	// You may also want to control width when formatting strings,
	// especially to ensure that they align in table-like output.
	// For basic right-justified width.
	pf("width4: |%6s|%6s|\n", "foo", "f")

	// To left-justify use the - flag as with numbers.
	pf("width5: |%-6s|%-6s|\n", "foo", "f")

	// So far we’ve seen Printf, which prints the formatted string to os.Stdout.
	// Sprintf formats and returns a string without printing it anywhere.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
