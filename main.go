package main

import (
	"fmt"
	"os"
)

func main() {
	println("Hello Go!")

	for i := 1; i < len(os.Args); i++ {
		println(os.Args[i])
	}
}

func println(message string) {
	fmt.Println(message)
}
