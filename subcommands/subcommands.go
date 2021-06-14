//* Lol wat?
package subcommands

//* Import all the things
import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"nebula/subcommands/installers"
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

			//* Run the installer
			installers.AppInstaller(c, packageName)

			//* or if it's a script
		} else if packageType == "script" {

			//* Run the installer
			installers.ScriptInstaller(c, packageName)

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

func println(message string) { fmt.Println(message) }

func Licence() {
	println("BSD-2-Clause Licence")
	println("")
	println("Copyright (c) 2021-Present Ruben Flores. All rights reserved.")
	println("")
	println("Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:")
	println("")
	println(" 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.")
	println("")
	println(" 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.")
	println("")
	println("THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
}
