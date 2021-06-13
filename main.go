package main

import (
	"fmt"
	"os"
)

func main() {
	for a := 1; a < len(os.Args); a++ {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			println("Welcome to the Brewery!")
			println("")
		}
	}
}

func println(message string) {
	fmt.Println(message)
}
