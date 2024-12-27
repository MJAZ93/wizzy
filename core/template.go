package core

import (
	"errors"
	"github.com/fatih/color"
	"wizzy/core/model"
	"wizzy/reader"
	"wizzy/writter"
)

func readTemplate(template model.Template, dir string) (model.Template, error) {
	var action Action

	if len(template.Folder) == 0 {
		action = Execute
	} else {
		a, err := navigateOrExecute(template)
		action = a
		if err != nil {
			return model.Template{}, errors.New("unable to decide if navigate or execute template at ./test. error:" + err.Error())
		}
	}

	if action == Navigate {
		option, err := navigateFolder(template)
		if err != nil {
			return model.Template{}, errors.New("unable to navigate to folder, error:" + err.Error())
		}

		t, err := reader.ReadTemplate(dir + "/" + option + "/")
		if err != nil {
			return model.Template{}, errors.New("unable to navigate to folder, error:" + err.Error())
		}

		return readTemplate(t, dir+"/"+option)
	} else if action == Execute {
		return template, nil
	} else {
		return model.Template{}, errors.New("invalid option")
	}
}

func processTemplate(template model.Template, dir string) error {
	color.Magenta("Processing template: " + dir + "/template.json")
	color.Yellow("Reading params")
	params, err := readParams(template)
	if err != nil {
		return errors.New("cant read params, err:" + err.Error())
	}

	//process rules
	color.Yellow("Processing rules")
	err = runRules(template, dir, params)
	if err != nil {
		return errors.New("cant run rules, err:" + err.Error())
	}

	return nil
}

func runRules(template model.Template, dir string, params []model.Param) error {
	for _, rule := range template.Rules {
		if rule.Rule == model.FileRule {

			if rule.MatchesRule(params) {
				// err := writter.CanWriteFile(rule, params)
				// if err != nil {
				//	return errors.New("verifying conditions for execution" + err.Error())
				//}

				err := writter.WriteFile(rule, params)
				if err != nil {
					return errors.New("executing rule execution" + err.Error())
				}
			}
		} else if rule.Rule == model.TemplateRule {
			t, err := reader.ReadTemplate(dir + "/" + rule.Destination + "/")
			if err != nil {
				return errors.New("unable to navigate to folder, error:" + err.Error())
			}

			return processTemplate(t, dir+"/"+rule.Destination)
		}
	}

	return nil
}
