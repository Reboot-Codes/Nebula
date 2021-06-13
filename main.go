package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "An alternative Brew client written in GoLang"
	app.Usage = ""
	app.Author = "Reboot-Codes"
	app.Version = "0.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "Install a Formulae or Cask (using --cask)",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "cask",
					Usage: "Tell the installer to look for a cask with that name",
				},
			},
			Action: func(c *cli.Context) {
				color.New(color.FgBlue).Printf("==>")
				fmt.Println(" Making a new instance of the Installer")
			},
		},
		{
			Name:  "search",
			Usage: "Search for Formulae or Casks",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "cask",
					Usage: "Filter for casks only",
				},
				&cli.BoolFlag{
					Name:  "formula",
					Usage: "Filter for formulae only",
				},
			},
			Action: func(c *cli.Context) {
				if c.NumFlags() == 0 {
					color.New(color.FgBlue).Printf("==>")
					fmt.Println(" Searching for \"" + c.Args()[0] + "\"")
				} else if c.NumFlags() == 1 {
					if c.Bool("cask") {
						color.New(color.FgBlue).Printf("==>")
						fmt.Println(" Searching for a cask called \"" + c.Args()[0] + "\"")
					} else if c.Bool("formula") {
						color.New(color.FgBlue).Printf("==>")
						fmt.Println(" Searching for a formula called \"" + c.Args()[0] + "\"")
					}
				} else if c.NumFlags() <= 2 {
					color.New(color.FgHiRed).Printf("==>")
					fmt.Println(" Error: Multiple Flags Specified")
					os.Exit(1)
				}
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
