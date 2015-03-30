package main

import (
	"github.com/docopt/docopt-go"
)

const usage = `Ninja Sphere Developer CLI.

Usage:
  ninja build <path>
  ninja -h | --help

Options:
  build         Builds a package for a driver or app, ready for distribution to the Ninja Sphere.
  -h --help     Show this screen.
`

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "Ninja Sphere Developer CLI", false)

	verifyDockerImages()

	if arguments["build"].(bool) {
		doBuild(arguments)
	}
}
