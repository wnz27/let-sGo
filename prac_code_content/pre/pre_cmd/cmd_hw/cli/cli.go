/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/10/1 21:30 10月
 **/
package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
	"sort"
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

func f5() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
// sort
func f6()  {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "Language for the greeting",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action:  func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action:  func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func f7()  {
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "language for the greeting",
				EnvVars: []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// filePath
func f8()  {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		&cli.StringFlag{
			Name: "password",
			Aliases: []string{"p"},
			Usage: "password for the mysql database",
			FilePath: "/etc/mysql/password",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/*
NAME:
Website Lookup CLI - Let's you query IPs, CNAMEs and Name Servers!

USAGE:
cli [global options] command [command options] [arguments...]

VERSION:
0.0.0

COMMANDS:
ns       Looks Up the NameServers for a Particular Host
cname    Looks up the CNAME for a particular host
ip       Looks up the IP addresses for a particular host
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h     show help
--version, -v  print the version
 */
func homework() {
	cmd := cli.App{
		Name: "Website Lookup CLI",
		Usage: "Let's you query IPs, CNAMEs and Name Servers!",
		Version: "0.0.0",

		Commands: []*cli.Command{
			{
				Name:    "ns",
				Usage:   "Looks Up the NameServers for a Particular Host",
				Action:  nsHandle,
			},
			{
				Name:    "cname",
				Usage:   "Looks up the CNAME for a particular host",
				Action:  cnameHandle,
			},
			{
				Name:    "ip",
				Usage:   "Looks up the IP addresses for a particular host",
				Action:  ipHandle,
			},
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getHost(c *cli.Context) string {
	return c.Args().First()
}

func ipHandle(c *cli.Context) error {
	host := getHost(c)
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return err
	}
	fmt.Println("The CNAME is:", cname)
	return nil
}

func cnameHandle(c *cli.Context) error {
	host := getHost(c)
	ips, err := net.LookupIP(host)
	if err != nil {
		return err
	}
	fmt.Println("ips are:")
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
	return nil
}

func nsHandle(c *cli.Context) error {
	host := getHost(c)
	ns, err := net.LookupNS(host)
	if err != nil {
		return err
	}
	fmt.Println("NSs are:")
	for _, v := range ns {
		fmt.Println(v.Host)
	}
	return nil
}



func main() {
	homework()
}


