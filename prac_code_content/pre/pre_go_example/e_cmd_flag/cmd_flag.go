/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/23 08:43 9月
 **/
package main

import (
	"flag"
	"fmt"
)

/*

Command-line flags are a common way to specify options for command-line programs.
For example, in wc -l the -l is a command-line flag.

Go provides a flag package supporting basic command-line flag parsing.
We’ll use this package to implement our example command-line program.
 */


func main() {
	// Basic flag declarations are available for string, integer, and boolean options.
	// Here we declare a string flag word with a default value "foo" and a short description.
	// This flag.String function returns a string pointer (not a string value);
	// we’ll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// This declares numb and fork flags, using a similar approach to the word flag.
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork",false, "an bool")

	// It’s also possible to declare an option that uses an existing var declared elsewhere in the program.
	// Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// Here we’ll just dump out the parsed options and any trailing positional arguments.
	// Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	/*
	To experiment with the command-line flags program it’s best to first compile it and
	then run the resulting binary directly.

	$ go build prac_code_content/pre/pre_go_example/e_cmd_flag/cmd_flag.go

	Try out the built program by first giving it values for all flags.
	$ ./cmd_flag -word=opt -numb=7 -fork -svar=flag
	output:
	word: opt
	numb: 7
	fork: true
	svar: flag
	tail: []

	Note that if you omit flags they automatically take their default values.
	$ ./cmd_flag -word=opt
	output:
	word: opt
	numb: 42
	fork: false
	svar: bar
	tail: []

	Trailing positional arguments can be provided after any flags.
	$ ./cmd_flag -word=opt a1 a2 a3
	output:
	word: opt
	numb: 42
	fork: false
	svar: bar
	tail: [a1 a2 a3]

	Note that the flag package requires all flags to appear before positional arguments
	(otherwise the flags will be interpreted as positional arguments).
	$ ./cmd_flag -word=opt a1 a2 a3 -numb=7
	output:
	word: opt
	numb: 42
	fork: false
	svar: bar
	tail: [a1 a2 a3 -numb=7]

	Use -h or --help flags to get automatically generated help text for the command-line program.
	$ ./cmd_flag -h
	output:
	Usage of ./cmd_flag:
	  -fork
	        an bool
	  -numb int
	        an int (default 42)
	  -svar string
	        a string var (default "bar")
	  -word string
	        a string (default "foo")

	If you provide a flag that wasn’t specified to the flag package,
	the program will print an error message and show the help text again.
	$ ./cmd_flag -wat
	output:
	flag provided but not defined: -wat
	Usage of ./cmd_flag:
	  -fork
	        an bool
	  -numb int
	        an int (default 42)
	  -svar string
	        a string var (default "bar")
	  -word string
	        a string (default "foo")
	*/
}
