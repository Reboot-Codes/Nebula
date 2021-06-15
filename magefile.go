// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Update

func printStatus(statusColor color.Attribute, statusMessage string) {

	color.New(statusColor).Printf("==>")
	fmt.Println(" " + statusMessage)

}

// A build step that requires additional params, or platform specific steps for example
func Build() error {

	mg.Deps(InstallDeps)
	printStatus(color.FgHiBlue, "Building...")
	cmd := exec.Command("go", "build", "-o", "nebula", ".")
	return cmd.Run()

}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	printStatus(color.FgHiBlue, "Installing...")
	return os.Rename("nebula", "/Users/reboot/go/bin/nebula")
}

// Update the binary
func Update() {
	mg.Deps(Uninstall)
	mg.Deps(Install)
}

// Custom uninstall script
func Uninstall() error {
	printStatus(color.FgHiBlue, "Uninstalling...")
	return os.RemoveAll("/Users/reboot/go/bin/nebula")
}

// Manage your deps, or running package managers.
func InstallDeps() {
	printStatus(color.FgHiBlue, "Installing Deps...")
	exec.Command("go", "get", "github.com/fatih/color").Run()
	exec.Command("go", "get", "github.com/urfave/cli").Run()
}

// Clean up after yourself
func Clean() {
	printStatus(color.FgHiBlue, "Cleaning...")
	os.RemoveAll("nebula")
}
