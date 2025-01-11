package core

import (
	"errors"
	"strings"
	"wizzy/core/model"
	ui_options "wizzy/core/ui/options"
	ui_textarea "wizzy/core/ui/textarea"
	ui_textinput "wizzy/core/ui/textinput"
)

func readParams(template model.Template, existingParams []model.Param) ([]model.Param, error) {
	var params []model.Param

	for _, f := range template.Parameters {
		parts := strings.SplitN(f.Condition, "==", 2)

		conditionalParam := ""
		conditionalParamValue := ""
		if len(parts) == 2 {
			conditionalParam = parts[0]
			conditionalParamValue = parts[1]
		}

		var data string

		existing := false
		conditionMet := true
		for _, ep := range existingParams {
			if f.ID == ep.Id {
				params = append(params, model.Param{
					Id:    ep.Id,
					Value: ep.Value,
				})
				existing = true
			}
		}

		if conditionalParam != "" && conditionalParamValue != "" {
			for _, ep := range append(params, existingParams...) {
				if ep.Id == conditionalParam {
					if ep.Value != conditionalParamValue {
						conditionMet = false
					}
				}
			}
		}

		if !existing && conditionMet {
			if f.Type == model.FreeType {
				text, err := ui_textinput.ReadText(f.Desc, f.Regex)
				if err != nil {
					return nil, errors.New("cant read param from text, err:" + err.Error())
				}
				data = text
			} else if f.Type == model.ListType {

			} else if f.Type == model.SelectType {
				err, choice := ui_options.GetOption(f.Desc, f.Options)
				if err != nil {
					return nil, errors.New("cant read param from select, err:" + err.Error())
				}
				data = choice
			} else if f.Type == model.FormattedType {
				text, err := ui_textarea.ReadText(f.Desc)
				if err != nil {
					return nil, errors.New("cant read param from text, err:" + err.Error())
				}
				data = text
			}

			params = append(params, model.Param{
				Id:    f.ID,
				Value: data,
			})
		}

		if !conditionMet {
			params = append(params, model.Param{
				Id:    f.ID,
				Value: "",
			})
		}
	}

	return params, nil
}
