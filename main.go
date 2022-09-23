package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `A simple command-line utility to list all cronjobs that are due between two timestamps.

Usage:
  crongap
  crongap -h | --help
  crongap --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/crongap v0.0.1")
	fmt.Println(opts)
}
