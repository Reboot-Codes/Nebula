//* Lol wat?
package subcommands

//* Import all the things
import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func Install(c *cli.Context) {

	// TODO: Acctually check if the package exists
	var foundPackage bool = true

	//* Default to installing a script,
	// TODO: Make this configurable
	var packageType string = "script"

	//* Get the name of the package
	var packageName string = c.Args()[0]

	//* Check if the --app flag is set
	if c.Bool("app") {

		//* Set the type of package to an app
		packageType = "app"

		//* Check if the --script flag is set
	} else if c.Bool("script") {

		//* Set the type of package to a script
		packageType = "script"

	}

	//* Check if the package exists
	if foundPackage {

		//* Check if the package is an app
		if packageType == "app" {

			//* Tell the user that they are installing a app with the name they passed in
			color.New(color.FgHiBlue).Printf("==>")
			fmt.Println(" Installing an App called \"" + packageName + "\"")

			//* or if it's a script
		} else if packageType == "script" {

			//* Tell the user that they are installing a script with the name that they passed in
			color.New(color.FgHiBlue).Printf("==>")
			fmt.Println(" Installing a Script called \"" + packageName + "\"")

		}

		//* if the package doesn't exist
	} else if !foundPackage {

		//* Tell the user what happened
		color.New(color.FgHiRed).Printf("Error:")
		fmt.Println(" Could not find any package called \"" + packageName + "\"")

		//* Exit with a non-zero exit code
		os.Exit(1)

	}

}

func Search(c *cli.Context) {

	//* Check the number of flags
	if c.NumFlags() == 0 {

		//* Search all packages
		color.New(color.FgHiBlue).Printf("==>")
		fmt.Println(" Searching for \"" + c.Args()[0] + "\"")

		//* Check if the user specified a search type
	} else if c.NumFlags() == 1 {

		//* Look for apps with that name
		if c.Bool("app") {

			//* Feedback
			color.New(color.FgHiBlue).Printf("==>")
			fmt.Println(" Searching for Apps")

			//* Look for planete with that name
		} else if c.Bool("script") {

			//* Feedback
			color.New(color.FgHiBlue).Printf("==>")
			fmt.Println(" Searching for Scripts")

		}

		//* Too many flags
	} else if c.NumFlags() <= 2 {

		//* Then don't specify a type >:(
		color.New(color.FgHiRed).Printf("Error:")
		fmt.Println(" Multiple Flags Specified")

		//* Exit with a non-zero exit code
		os.Exit(1)

	}

}
