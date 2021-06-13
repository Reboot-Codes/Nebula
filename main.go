//* Lol wat?
package main

//* Import all the things
import (
	"log"
	"os"

	"github.com/urfave/cli"

	"brewery/subcommands"
)

//* Instantiate a new app
var app = cli.NewApp()

//* All about the CLI
func info() {
	app.Name = "An alternative Brew client written in GoLang"
	app.Usage = ""
	app.Author = "Reboot-Codes"
	app.Version = "0.1.0"
}

//* Specify ALL the things that make it all work
func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "Install a Formulae (using --formula) or Cask (using --cask)",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "cask",
					Usage: "Tell the installer to install a cask with that name",
				},
				&cli.BoolFlag{
					Name:  "formula",
					Usage: "Tell the installer to install a formula with that name",
				},
			},
			Action: func(c *cli.Context) {
				subcommands.Install(c)
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
				subcommands.Search(c)
			},
		},
	}
}

//* The rEaL function that makes it all work
func main() {

	//* Set the info for our CLI
	info()
	//* Set the command descriptions, flags, options, and actions of the commands for the CLI
	commands()

	//* Handle errors by printing to the console
	//? No idea how this works tho...
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
