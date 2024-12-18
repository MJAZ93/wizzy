package core

import (
	ui_options "wizzy/core/ui/options"
	ui_textinput "wizzy/core/ui/textinput"
)

//read template (path)
//run
//read params

func Run() {
	err, choice := ui_options.GetOption("Please choose an option (Enter to select, q to cancel)", []string{"trash", "recycle", "can", "mamob"})
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Selected option:", choice)
	}

	text, err := ui_textinput.ReadText("Type any")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Text typed:", text)
	}

}

func runRules(rule string) {}
func readParams()          {}
