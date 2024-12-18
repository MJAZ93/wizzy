package reader

import (
	"wizzy/core/model"
)

func ReadTemplate(filePath string) (model.Template, error) {
	return readTemplate(filePath)
}
