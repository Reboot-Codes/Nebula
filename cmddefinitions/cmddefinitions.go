package cmddefinitions

import (
	"nebula/subcommands"

	"github.com/urfave/cli"
)

//* Specify ALL the things that make it all work
func Commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:  "search",
			Usage: "Search for packages",
			Action: func(c *cli.Context) {
				subcommands.Search(c)
			},
		},
		{
			Name:  "install",
			Usage: "Install a package",
			Action: func(c *cli.Context) {
				subcommands.Install(c)
			},
		},
		{
			Name:  "remove",
			Usage: "Removes an installed package",
			Action: func(c *cli.Context) {
				subcommands.Remove(c)
			},
		},
		{
			Name:  "licence",
			Usage: "Prints the licence text to the console",
			Action: func(c *cli.Context) {
				subcommands.Licence()
			},
		},
	}
}
