package main

import (
	"github.com/fatih/color"
	"wizzy/core"
)

// todo: edit logic
// core logic
func main() {
	cyan := color.New(color.FgCyan)

	boldCyan := cyan.Add(color.Bold).Add(color.Underline)
	boldCyan.Println("Starting Wizzy")

	color.Cyan("working folder ./test")

	core.Run("test")
}
