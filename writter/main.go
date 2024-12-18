package writter

import (
	"errors"
	"wizzy/core/model"
)

func WriteFile(rule model.Rule, params []model.Param) error {
	//verify file destination
	fileDestinationPath := replacePlaceholders(rule.Destination, params)
	fileOriginPath := replacePlaceholders(rule.Origin, params)
	fileExists, err := fileExistsInDir(fileDestinationPath, fileOriginPath)
	if err != nil {
		return errors.New("will be unable to verify dir: '" + rule.Path + "'" + "'. error: " + err.Error())
	}

	//verify template file destination
	templateFilePath := rule.Path + "/" + rule.Origin
	templateOrigin := ""
	if fileExists {
		templateFilePath = templateFilePath + ".e"
		templateOrigin = rule.Origin + ".e"
	} else {
		templateFilePath = templateFilePath + ".n"
		templateOrigin = rule.Origin + ".n"
	}
	templateExist, err := fileExistsInDir(rule.Path, templateOrigin)
	if err != nil {
		return errors.New("will be unable to verify dir: '" + rule.Path + "'" + "'. error: " + err.Error())
	}

	if !templateExist {
		return errors.New("template: '" + templateFilePath + "' does not exists")
	}

	if fileExists {
		//edit logic
		err = insertProcessedBlocks(templateFilePath, fileDestinationPath+"/"+fileOriginPath, params)
		if err != nil {
			return errors.New("unable to add templated content to file in the dir:'" + fileDestinationPath + "/" + fileOriginPath + "'. error: " + err.Error())
		}
	} else {
		content, err := generateFileContent(templateFilePath, params)
		if err != nil {
			return errors.New("unable to generate templated file in the dir:'" + rule.Path + "'. error: " + err.Error())
		}

		_, err = createFileWithContent(fileDestinationPath, fileOriginPath, content)
		if err != nil {
			return errors.New("unable to create files in the dir:'" + fileDestinationPath + "'. error: " + err.Error())
		}
	}

	return nil
}

func CanWriteFile(rule model.Rule, params []model.Param) error {
	//todo: check if file and template exists and match
	// check if .n file does not contains any @@
	// check if destination files has all @@ that .e has
	err := canCreateFileInDir(rule.Path)
	if err != nil {
		return errors.New("will be unable to create files in the dir: '" + rule.Path + "'" + "'. error: " + err.Error())
	}

	exist, err := fileExistsInDir(rule.Path, rule.Origin)
	if err != nil {
		return errors.New("will be unable to verify dir: '" + rule.Path + "'" + "'. error: " + err.Error())
	}

	templateFilePath := rule.Path + "/" + rule.Origin
	if exist {
		templateFilePath = templateFilePath + ".e"
	} else {
		templateFilePath = templateFilePath + ".n"
	}

	_, err = checkTemplateParams(params, templateFilePath)
	if err != nil {
		return errors.New("will be unable to parse template: '" + rule.Path + "/" + rule.Origin + "'" + "'. error: " + err.Error())
	}

	return nil
}
