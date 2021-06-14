package installers

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func AppInstaller(c *cli.Context, packageName string) {

	//* Tell the user that they are installing a app with the name they passed in
	color.New(color.FgHiBlue).Printf("==>")
	fmt.Println(" Installing an App called \"" + packageName + "\"")

}

func ScriptInstaller(c *cli.Context, packageName string) {

	//* Tell the user that they are installing a script with the name that they passed in
	color.New(color.FgHiBlue).Printf("==>")
	fmt.Println(" Installing a Script called \"" + packageName + "\"")

}