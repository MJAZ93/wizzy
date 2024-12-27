package core

import (
	"fmt"
	"wizzy/core/model"
	ui_options "wizzy/core/ui/options"
)

func navigateOrExecute(template model.Template) (Action, error) {
	title := fmt.Sprintf("Template: %s:", template.TemplateDesc.Name)

	err, choice := ui_options.GetOption(title, []string{Navigate.String(), Execute.String()})

	return ActionFromString(choice), err
}

func navigateFolder(template model.Template) (string, error) {
	title := "Navigate to:"

	var options []string
	for _, f := range template.Folder {
		options = append(options, f.Name)
	}

	err, choice := ui_options.GetOption(title, options)

	return choice, err
}
