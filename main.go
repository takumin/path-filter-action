package main

import (
	"os"

	"github.com/takumin/path-filter-action/internal/command"
)

var osExit = os.Exit

func main() {
	osExit(command.Main(
		os.Stdout,
		os.Stderr,
		os.Stdin,
		os.Args,
	))
}
