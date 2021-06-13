//* Lol wat?
package main

//* Import all the things
import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
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
				// TODO: Acctually check if the package exists
				var foundPackage bool = true
				//* Default to installing a formula,
				// TODO: Make this configurable
				var packageType string = "formula"
				//* Get the name of the package
				var packageName string = c.Args()[0]

				//* Check if the --cask flag is set
				if c.Bool("cask") {
					//* Set the type of package to a cask
					packageType = "cask"
					//* Check if the --formula flag is set
				} else if c.Bool("formula") {
					//* Set the type of package to a formula
					packageType = "formula"
				}

				//* Check if the package exists
				if foundPackage {
					//* Check if the package is a cask
					if packageType == "cask" {
						//* Tell the user that they are installing a cask with the name they passed in
						color.New(color.FgHiBlue).Printf("==>")
						fmt.Println(" Installing a cask called \"" + packageName + "\"")
						//* or if it's a formula
					} else if packageType == "formula" {
						//* Tell the user that they are installing a formula with the name that they passed in
						color.New(color.FgHiBlue).Printf("==>")
						fmt.Println(" Installing a formula called \"" + packageName + "\"")
					}
					//* if the package doesn't exist
				} else if !foundPackage {
					//* Tell the user what happened
					color.New(color.FgHiRed).Printf("Error:")
					fmt.Println(" Could not find any package called \"" + packageName + "\"")
					//* Exit with a non-zero exit code
					os.Exit(1)
				}
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
				//* Check the number of flags
				if c.NumFlags() == 0 {
					//* Search all packages
					color.New(color.FgHiBlue).Printf("==>")
					fmt.Println(" Searching for \"" + c.Args()[0] + "\"")
					//* Check if the user specified a search type
				} else if c.NumFlags() == 1 {
					//* Look for casks with that name
					if c.Bool("cask") {
						color.New(color.FgHiBlue).Printf("==>")
						fmt.Println(" Searching for a cask called \"" + c.Args()[0] + "\"")
						//* Look for formulae with that name
					} else if c.Bool("formula") {
						color.New(color.FgHiBlue).Printf("==>")
						fmt.Println(" Searching for a formula called \"" + c.Args()[0] + "\"")
					}
				} else if c.NumFlags() <= 2 {
					//* Then don't specify a type >:(
					color.New(color.FgHiRed).Printf("Error:")
					fmt.Println(" Multiple Flags Specified")
					//* Exit with a non-zero exit code
					os.Exit(1)
				}
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
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
