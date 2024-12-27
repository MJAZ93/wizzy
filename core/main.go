package core

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	ui_progress "wizzy/core/ui/progress"
	ui_textarea "wizzy/core/ui/textarea"
	ui_textinput "wizzy/core/ui/textinput"
	"wizzy/reader"
)

// read template (path)
// run
// read params
func Run(folder string) {
	fmt.Println("Reading template:" + "./" + folder + "/template.json")

	template, err := reader.ReadTemplate("./" + folder + "/")
	if err != nil {
		color.Red("error reading template, err: " + err.Error())
		log.Fatalf("exit with error: %v", err)
		return
	}

	finalTemplate, err := readTemplate(template, "./"+folder)
	if err != nil {
		color.Red("error reading template, err: " + err.Error())
		log.Fatalf("exit with error: %v", err)
		return
	}

	err = processTemplate(finalTemplate, "./"+folder)
	if err != nil {
		fmt.Println("err:", err.Error())
		log.Fatalf("exit with error: %v", err)
		return
	}

	color.Green("Successfully processed all rules")
}

func test() {
	ui_progress.Show()

	text, err := ui_textinput.ReadText("Type any", "")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Text typed:", text)
	}

	text, err = ui_textarea.ReadText("Text area title")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Text typed:", text)
	}
}
