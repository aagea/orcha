package main

import (
	"github.com/aagea/orcha/cmd/orcha/commands"
)

// Version of the command
var Version string

// Commit from which the command was built
var Commit string

func main() {
	commands.Execute(Version, Commit)
}
