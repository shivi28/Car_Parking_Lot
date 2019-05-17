package main

import (
	"flag"
	"fmt"
)

func main() {

	fmt.Println("app started ....")

	flag.Parse()

	if len(flag.Args()) > 0 {
		// To run code from input file
		RunFileCommand(flag.Args()[0])
		return
	}

	// To run code from command line
	RunCliCommand()

}
