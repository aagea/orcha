package main

import "fmt"

// Version of the command
var Version string

// Commit from which the command was built
var Commit string

func main() {
	fmt.Printf("ORCHA %s(%s)", Version, Commit)
}
