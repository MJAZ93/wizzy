package core

import (
	"errors"
	"wizzy/core/model"
	ui_options "wizzy/core/ui/options"
	ui_textarea "wizzy/core/ui/textarea"
	ui_textinput "wizzy/core/ui/textinput"
)

func readParams(template model.Template) ([]model.Param, error) {
	var params []model.Param

	for _, f := range template.Parameters {
		var data string

		//todo: validate regex
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

	return params, nil
}
