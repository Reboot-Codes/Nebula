package installers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

//* Helper function to print the current status to the console
func printStatus(statusColor color.Attribute, statusMessage string) {

	color.New(statusColor).Printf("==>")
	fmt.Println(" " + statusMessage)

}

//* Installer for all packages
func AppInstaller(c *cli.Context, packageName string) {

	//* Tell the user that they are installing a app with the name they passed in
	printStatus(color.FgHiBlue, "Getting Manifest Data for: "+packageName)

	//* Get the namifest data for the package if it was found
	// TODO: Pull the real manifest data from the repo, but for now, use a hard coded var
	var foundManifestData bool = true

	//* If we did find the manifest data for the specifed package
	if foundManifestData {

		//* Get the environment variable of what processor we're running on
		var processorType string = os.Getenv("NEBULA_PROCESSOR_ARCH")

		// TODO: Really check the compatibility of the package
		var packageCompatible bool = true

		printStatus(color.FgHiGreen, "Found Manifest Data for: "+packageName)

		//* Looks like you forgot something...
		if processorType == "" {

			//* Give an error if the type of processor isn't specified
			printStatus(color.FgHiRed, "Error: Your processor architecture has not been set!")
			printStatus(color.FgHiRed, "For arm macs, run: export NEBULA_PROCESSOR_ARCH=arm64")
			printStatus(color.FgHiRed, "or")
			printStatus(color.FgHiRed, "For Intel macs, run: export NEBULA_PROCESSOR_ARCH=amd64")

			//* Boomer
		} else if processorType == "amd64" {

			//* Continue installing the package
			printStatus(color.FgHiBlue, "Getting Binary for the platform \"amd64\" for the package \""+packageName+"\"")

			//* "The future of mac" or something like that
		} else if processorType == "arm64" {

			//* Check the direct compatibility of the package
			if packageCompatible {

				//* Continue installing the package
				printStatus(color.FgHiBlue, "Getting Binary for the platform \"arm64\" for the package \""+packageName+"\"")

				//* I guess it isn't directly compatible
			} else if !packageCompatible {

				//* Warn about possible errors installing under Rosseta 2
				printStatus(color.FgHiYellow, "Warn: Base binary incompatible with the platform \"arm64\" for the package \""+packageName+"\"")
				printStatus(color.FgHiYellow, "Warn: Installing \""+packageName+"\" using Rosseta 2 may cause errors!")

				//* Continue installing the package
				printStatus(color.FgHiBlue, "Getting Binary for the platform \"arm64\" for the package \""+packageName+"\"")

			}

		}

		//* Couldn't find the manifest data
	} else if !foundManifestData {

		//* Oops, that's an error
		printStatus(color.FgHiRed, "Error: Couldn't Find the Manifest Data for \""+packageName+"\"")
		printStatus(color.FgHiRed, "Check if you spelled the name of the package correctly, ")
		printStatus(color.FgHiRed, "becuase we couldn't find a package with that name in any manifest...")

	}

}

//* Remove apps because you need that in a package manager
func AppUnInstaller(c *cli.Context, packageName string) {

	//* RUN THE UNINSTALLER!
	printStatus(color.FgHiBlue, "Running Unistaller for \""+packageName+"\"...")

}
