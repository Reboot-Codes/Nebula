//* Lol wat?
package main

//* Import all the things
import (
	"log"
	"os"

	"github.com/urfave/cli"

	"brewery/cmddefinitions"
)

//* Instantiate a new app
var App = cli.NewApp()

//* All about the CLI
func info() {
	App.Name = "An alternative Brew client written in GoLang"
	App.Usage = ""
	App.Author = "Reboot-Codes"
	App.Version = "0.1.0"
}

//* The rEaL function that makes it all work
func main() {

	//* Set the info for our CLI
	info()
	//* Set the command descriptions, flags, options, and actions of the commands for the CLI
	cmddefinitions.Commands(App)

	//* Handle errors by printing to the console
	//? No idea how this works tho...
	err := App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
