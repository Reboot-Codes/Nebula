package cmddefinitions

import (
	"brewery/subcommands"

	"github.com/urfave/cli"
)

//* Specify ALL the things that make it all work
func Commands(app *cli.App) {
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
