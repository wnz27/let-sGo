/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/22 5:10 下午
* Description:
 */
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Join should be used to construct paths in a portable way.
	// It takes any number of arguments and constructs a hierarchical path from them.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// You should always use Join instead of concatenating /s or \s manually.
	// In addition to providing portability, Join will also normalize paths by removing superfluous separators and
	// directory changes.
	fmt.Println("remove superfluous separators :", filepath.Join("dir1//", "filename1"))
	fmt.Println("remove superfluous directory :", filepath.Join("dir1/../dir1", "filename2"))

	// Dir and Base can be used to split a path to the directory and the file.
	// Alternatively, Split will return both in the same call.
	fmt.Println("Dir(p): ", filepath.Dir(p))
	fmt.Println("Base(p): ", filepath.Base(p))
	d1, b1 := filepath.Split(p)
	fmt.Println("Split(p): ", d1, "|", b1)

	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))  // false
	fmt.Println(filepath.IsAbs("/dir/file2"))  // true

	filename := "config.json"

	// Some file names have extensions following a dot.
	// We can split the extension out of such names with Ext.
	ext := filepath.Ext(filename)
	fmt.Println(ext)  // .json

	// To find the file’s name with the extension removed, use strings.TrimSuffix.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target.
	// It returns an error if the target cannot be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
