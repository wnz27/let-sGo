/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/10/1 21:30 10月
 **/
package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func f1() {
	app := &cli.App{
		Name: "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// use Arguments
func f2() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Flags
func f3()  {
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// You can also set a destination variable for a flag, to which the content will be scanned.
func f4() {
	var language string

	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "someone"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f4()
}


