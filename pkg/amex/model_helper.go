package amex

func (amex *AmexEngine) createModelFiles() {
	modelDir := amex.microServicDirStructure["models"].(map[string]interface{})["dir"].(string)
	modelPackageName := amex.microServicDirStructure["models"].(map[string]interface{})["name"].(string)

	for _, object := range amex.microServiceData.Spec.Objects {
		modelGoFileName := strToGoFileName(object.Singular)
		modelName := strToStructName(object.Singular)

		modelPath := modelDir + "/" + modelGoFileName
		createFile(modelPath)

		strModel := "package " + modelPackageName + "\n\n"
		strModel = strModel + "type " + modelName + " struct {\n"
		//strModel = strModel + "\tId string `yaml:\"id\"`\n"

		strModel = strModel + "\n\t//Persistent attributes\n"
		for _, att := range object.Attributes.Persistent {
			attName := strToStructAttName(att.Name)
			attJson := strToStructAttJSON(att.Name)
			strModel = strModel + "\t" + attName + " " + att.PersistentSpec.DataType + " `json:\"" + attJson + "\"`\n"

		}

		strModel = strModel + "\n\t//Relational attributes\n"
		for _, att := range object.Attributes.Relational {
			attName := strToStructAttName(att.Name)
			attJson := strToStructAttJSON(att.Name)

			if att.Type == "1:N" || att.Type == "N:N" {
				strModel = strModel + "\t" + attName + " []" + strToStructName(att.Object) + " `json:\"" + attJson + "\"`\n"

			} else if att.Type == "1:1" {
				strModel = strModel + "\t" + attName + " " + strToStructName(att.Object) + " `json:\"" + attJson + "\"`\n"
			}
		}

		strModel = strModel + "}"
		appendFile(modelPath, strModel)

	}
}
