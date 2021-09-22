/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/22 5:55 下午
* Description:
 */
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

/*
	Go has several useful functions for working with directories in the file system.
 */

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printFileMode() {
	fmt.Println(fs.ModeDir)
	fmt.Println(fs.ModeAppend)
	fmt.Println(fs.ModeCharDevice)
	fmt.Println(fs.ModeDevice)
	fmt.Println(fs.ModeExclusive)
	fmt.Println(fs.ModeSocket)
	fmt.Println(fs.ModeType)
	fmt.Println(fs.ModeTemporary)
}


func main() {
	os.RemoveAll("subdir")
	//printFileMode()
	// Create a new sub-directory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// When creating temporary directories, it’s good practice to defer their removal.
	// os.RemoveAll will delete a whole directory tree (similarly to rm -rf).
	defer os.RemoveAll("subdir")

	// Helper function to create a new empty file.
	creatEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	creatEmptyFile("subdir/file1")

	// We can create a hierarchy of directories, including parents with MkdirAll.
	// This is similar to the command-line mkdir -p.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	creatEmptyFile("subdir/parent/file2")
	creatEmptyFile("subdir/parent/file3")
	creatEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	count := 0
	for _, entry := range c {
		fI, _ := entry.Info()
		fmt.Println(count, ":", entry.Name(), entry.IsDir(), entry.Type(), fI.Mode())
		count += 1
	}

	// Chdir lets us change the current working directory, similarly to cd.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Now we’ll see the contents of subdir/parent/child when listing the current directory.
	c, err = os.ReadDir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started.
	err = os.Chdir("../../")
	check(err)
	c, err = os.ReadDir(".")
	fmt.Println("Listing ../../")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	fmt.Println("Visiting subdir")
	// We can also visit a directory recursively, including all its sub-directories.
	// Walk accepts a callback function to handle every file or directory visited.
	err = filepath.Walk("subdir", visit)

}

// visit is called for every file or directory found recursively by filepath.Walk.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
