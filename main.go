package main

import (
	"github.com/fatih/color"
	"os"
	"wizzy/core"
)

// todo: edit logic
// core logic
func main() {
	cyan := color.New(color.FgCyan)

	boldCyan := cyan.Add(color.Bold).Add(color.Underline)
	boldCyan.Println("Starting Wizzy")

	folder := "./"

	args := os.Args
	if len(args) < 2 {
		folder = folder + ".wizzy"
	} else {
		folder = folder + ".wizzy/" + args[1]
	}

	//todo: add clear descriptions of the params to not confuse

	color.Cyan("working folder " + folder)

	core.Run(folder)
}
