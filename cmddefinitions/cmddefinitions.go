package cmddefinitions

import (
	"nebula/subcommands"

	"github.com/urfave/cli"
)

//* Specify ALL the things that make it all work
func Commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "Install a script (using --script) or an app (using --app)",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "app",
					Usage: "Tell the installer to install an app with that name",
				},
				&cli.BoolFlag{
					Name:  "script",
					Usage: "Tell the installer to install a script with that name",
				},
			},
			Action: func(c *cli.Context) {
				subcommands.Install(c)
			},
		},
		{
			Name:  "search",
			Usage: "Search for scripts or apps",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "app",
					Usage: "Filter for apps only",
				},
				&cli.BoolFlag{
					Name:  "script",
					Usage: "Filter for scripts only",
				},
			},
			Action: func(c *cli.Context) {
				subcommands.Search(c)
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
