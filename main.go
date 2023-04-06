package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		os.Exit(1)
	}
}

func run() {
	fmt.Printf("Running: %v\n", os.Args[2:])
}
