package amex

func (amex *AmexEngine) createMicroServiceFileStructure() {

	for _, dir := range amex.microServicDirStructure {
		dirPath := dir.(map[string]interface{})["dir"].(string)
		createDir(dirPath)

		fileName := dir.(map[string]interface{})["file"].(string)
		if fileName != "" {
			fileName = dirPath + "/" + fileName
			createFile(fileName)

			packageName := dir.(map[string]interface{})["name"].(string)
			if packageName != "" {
				packageName = "package " + packageName + "\n\n"
				appendFile(fileName, packageName)
			}
		}
	}
}

/*
func setupManyToManyDBRelationship(amex *Amex) {
	joinTableList := []string{}
	for _, entity := range amex.serviceData.Spec.Entities {

		for _, att := range entity.Attributes {
			if strings.HasPrefix(att.Attspec.Type, "entity") {
				relEntity := amex.getEntityByName(strings.Split(att.Attspec.Type, ".")[1])
				if att.Attspec.Length == "*" {
					if !amex.notManyToMany(relEntity, entity) {
						entityExists := false
						relEntityExists := false

						for _, str := range joinTableList {
							if str == entity.Name {
								entityExists = true
							}
							if str == entity.Name {
								relEntityExists = true
							}
						}

						if !entityExists && !relEntityExists {
							amex.setupJoinTable(relEntity, entity)
							joinTableList = append(joinTableList, entity.Name)
							joinTableList = append(joinTableList, relEntity.Name)
						}
					}
				}

			}
		}
	}
}

func (amex *Amex) setupJoinTable(relEntity, entity Entity) {
	schemaDir := amex.serviceDirStructure["schema"].(map[string]interface{})["dir"].(string)

	dbFileName := strToDbFileName("junction_" + entity.Singular + "_with_" + relEntity.Singular)
	dbTableName := strToDbTableName("junction_" + entity.Plural + "_with_" + relEntity.Plural)
	// db file
	dbFilePath := schemaDir + "/" + dbFileName
	createFile(dbFilePath)

	dbSchemaName := strings.ToUpper(amex.serviceData.Metadata.Name)
	strDB := "CREATE SCHEMA IF NOT EXISTS `" + dbSchemaName + "` DEFAULT CHARACTER SET utf8;\n\n"
	strDB = strDB + "USE `" + dbSchemaName + "`;\n\n"

	strDB = strDB + "DROP TABLE IF EXISTS `" + dbSchemaName + "`.`" + dbTableName + "`;\n\n"
	strDB = strDB + "CREATE TABLE `" + dbSchemaName + "`.`" + dbTableName + "` (\n"
	strDB = strDB + "  `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,\n"

	strDB = strDB + "\n  -- Metadata columns --\n"
	for _, matt := range amex.dbMetadataAtt.Attributes {
		// db file
		dbColumn := "  `" + matt.DBName + "` " + matt.DataType
		if matt.Nullable != "" {
			dbColumn = dbColumn + " " + matt.Nullable
		}

		if matt.Defualt != "" {
			dbColumn = dbColumn + " DEFAULT " + matt.Defualt
		}
		dbColumn = dbColumn + ",\n"
		strDB = strDB + dbColumn

	}

	strDB = strDB + "\n  -- Relationship columns --\n"
	columnName := strToDbColumnName(entity.Singular + "_ID")
	strDB = strDB + "  `" + columnName + "` BIGINT UNSIGNED,\n"

	columnName = strToDbColumnName(relEntity.Singular + "_ID")
	strDB = strDB + "  `" + columnName + "` BIGINT UNSIGNED,\n"

	strDB = strDB + "\n  -- Constraint --\n"

	dbTablenameSingular := strToDbTableName(entity.Singular)

	tablename := strToDbTableName(entity.Plural)
	strDB = strDB + "  CONSTRAINT `FK_" + dbTablenameSingular + "` FOREIGN KEY (`" + dbTablenameSingular + "_ID`) REFERENCES `" + dbSchemaName + "`.`" + tablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

	dbTablenameSingular = strToDbTableName(relEntity.Singular)

	tablename = strToDbTableName(relEntity.Plural)
	strDB = strDB + "  CONSTRAINT `FK_" + dbTablenameSingular + "` FOREIGN KEY (`" + dbTablenameSingular + "_ID`) REFERENCES `" + dbSchemaName + "`.`" + tablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

	strDB = strDB + "  PRIMARY KEY (`ID`) USING BTREE"
	strDB = strDB + "\n) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;\n"
	appendFile(dbFilePath, strDB)

}

func (amex *Amex) notManyToMany(relEntity, entity Entity) bool {
	for _, e := range amex.serviceData.Spec.Entities {
		if e.Name == relEntity.Name {
			for _, att := range e.Attributes {
				if strings.HasPrefix(att.Attspec.Type, "entity") {
					if strings.Split(att.Attspec.Type, ".")[1] == entity.Name {
						return false
					}
				}
			}
		}
	}
	return true
}

func (amex *Amex) setupDbTableConstraintAndRelationshipAttWithEndStatement() {
	schemaDir := amex.serviceDirStructure["schema"].(map[string]interface{})["dir"].(string)

	for _, entity := range amex.serviceData.Spec.Entities {
		dbFileName := strToDbFileName(entity.Singular)
		dbFilePath := schemaDir + "/" + dbFileName
		strDB := "\n  -- Relationship columns --\n"
		appendFile(dbFilePath, strDB)
	}
	// Relationships attr
	for _, entity := range amex.serviceData.Spec.Entities {

		for _, att := range entity.Attributes {
			if strings.HasPrefix(att.Attspec.Type, "entity") {
				relEntity := amex.getEntityByName(strings.Split(att.Attspec.Type, ".")[1])
				dbFileName := strToDbFileName(relEntity.Singular)

				dbFilePath := schemaDir + "/" + dbFileName

				if strings.Split(att.Attspec.Type, ".")[1] == entity.Name {
					strDB := strToDbColumnName(att.Name + "_ID")
					strDB = "  `" + strDB + "` BIGINT UNSIGNED,\n"

					appendFile(dbFilePath, strDB)
				} else {

					if att.Attspec.Length == "1" {
						strDB := strToDbColumnName(entity.Singular + "_ID")
						strDB = "  `" + strDB + "` BIGINT UNSIGNED,\n"

						appendFile(dbFilePath, strDB)

					}
					if att.Attspec.Length == "*" {
						if amex.notManyToMany(relEntity, entity) {
							strDB := strToDbColumnName(entity.Singular + "_ID")
							strDB = "  `" + strDB + "` BIGINT UNSIGNED,\n"

							appendFile(dbFilePath, strDB)
						}
					}
				}
			}
		}
	}

	for _, entity := range amex.serviceData.Spec.Entities {
		dbFileName := strToDbFileName(entity.Singular)
		dbFilePath := schemaDir + "/" + dbFileName
		strDB := "\n  -- Constraint --\n"
		appendFile(dbFilePath, strDB)
	}
	// constraint FK attr
	dbSchemaName := strings.ToUpper(amex.serviceData.Metadata.Name)
	for _, entity := range amex.serviceData.Spec.Entities {

		for _, att := range entity.Attributes {
			if strings.HasPrefix(att.Attspec.Type, "entity") {

				relEntity := amex.getEntityByName(strings.Split(att.Attspec.Type, ".")[1])

				dbFileName := strToDbFileName(relEntity.Singular)
				dbTablename := strToDbTableName(entity.Plural)
				dbTablenameSingular := strToDbTableName(entity.Singular)
				dbFilePath := schemaDir + "/" + dbFileName

				if strings.Split(att.Attspec.Type, ".")[1] == entity.Name {
					refDBTablename := strToDbTableName(relEntity.Singular)
					strDB := "  CONSTRAINT `FK_" + dbTablenameSingular + "_" + strToDbTableName(att.Name) + "_" + refDBTablename + "` FOREIGN KEY (`" + strToDbTableName(att.Name) + "_ID`) REFERENCES `" + dbSchemaName + "`.`" + dbTablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

					appendFile(dbFilePath, strDB)
				} else {
					if att.Attspec.Length == "1" {
						refDBTablename := strToDbTableName(relEntity.Singular)
						strDB := "  CONSTRAINT `FK_" + dbTablenameSingular + "_" + refDBTablename + "` FOREIGN KEY (`" + dbTablenameSingular + "_ID`) REFERENCES `" + dbSchemaName + "`.`" + dbTablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

						appendFile(dbFilePath, strDB)

					}
					if att.Attspec.Length == "*" {
						if amex.notManyToMany(relEntity, entity) {
							refDBTablename := strToDbTableName(relEntity.Singular)
							strDB := "  CONSTRAINT `FK_" + dbTablenameSingular + "_" + refDBTablename + "` FOREIGN KEY (`" + dbTablenameSingular + "_ID`) REFERENCES `" + dbSchemaName + "`.`" + dbTablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

							appendFile(dbFilePath, strDB)
						}
					}
				}
			}
		}

	}

	// constraint PK attr
	for _, entity := range amex.serviceData.Spec.Entities {
		dbFileName := strToDbFileName(entity.Singular)
		dbFilePath := schemaDir + "/" + dbFileName

		strDB := "  PRIMARY KEY (`ID`) USING BTREE"

		appendFile(dbFilePath, strDB)
	}

	// End statement
	for _, entity := range amex.serviceData.Spec.Entities {
		dbFileName := strToDbFileName(entity.Singular)
		dbFilePath := schemaDir + "/" + dbFileName

		strDB := "\n) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;\n"
		appendFile(dbFilePath, strDB)
	}

}

func (amex *Amex) setupMedelRelationshipAtt() {
	modelDir := amex.serviceDirStructure["models"].(map[string]interface{})["dir"].(string)

	for _, entity := range amex.serviceData.Spec.Entities {
		modelGoFileName := strToGoFileName(entity.Singular)
		modelPath := modelDir + "/" + modelGoFileName
		strModel := ""

		for _, att := range entity.Attributes {
			if strings.HasPrefix(att.Attspec.Type, "entity") {

				relEntity := amex.getEntityByName(strings.Split(att.Attspec.Type, ".")[1])

				if att.Attspec.Length == "1" {
					attName := strToStructAttName(att.Name + "_E_Obj")
					attJson := strToStructAttJSON(att.Name + "_E_Obj")

					strModel = strModel + "\t" + attName + " " + strToStructName(relEntity.Singular) + " `json:\"" + attJson + "\"`\n"
				}
				if att.Attspec.Length == "*" {
					attName := strToStructAttName(att.Name + "_E_List")
					attJson := strToStructAttJSON(att.Name + "_E_List")

					strModel = strModel + "\t" + attName + " []" + strToStructName(relEntity.Singular) + " `json:\"" + attJson + "\"`\n"
				}

			}
		}
		if strModel != "" {
			strModel = "\n\t//Relationship attributes\n" + strModel
			strModel = strModel + "}"
			appendFile(modelPath, strModel)
		} else {
			strModel = "}"
			appendFile(modelPath, strModel)
		}

	}
}

func (amex *Amex) setupModelMainAtt() {
	modelDir := amex.serviceDirStructure["models"].(map[string]interface{})["dir"].(string)

	for _, entity := range amex.serviceData.Spec.Entities {

		modelGoFileName := strToGoFileName(entity.Singular)
		modelPath := modelDir + "/" + modelGoFileName

		strModel := "\n\t//Main attributes\n"

		for _, att := range entity.Attributes {
			if !strings.HasPrefix(att.Attspec.Type, "entity") {

				attName := strToStructAttName(att.Name)
				attJson := strToStructAttJSON(att.Name)
				strModel = strModel + "\t" + attName + " " + att.Attspec.Type + " `yaml:\"" + attJson + "\"`\n"

			}
		}
		appendFile(modelPath, strModel)
	}
}

func (amex *Amex) getEntityByName(str string) Entity {
	for _, entity := range amex.serviceData.Spec.Entities {
		if entity.Name == str {
			return entity
		}
	}
	return Entity{}
}

func (amex *Amex) setupModelMetadatAtt() {
	modelDir := amex.serviceDirStructure["models"].(map[string]interface{})["dir"].(string)

	for _, entity := range amex.serviceData.Spec.Entities {
		modelGoFileName := strToGoFileName(entity.Singular)
		modelPath := modelDir + "/" + modelGoFileName
		strModel := "\n\t//Metadata attributes\n"

		for _, matt := range amex.dbMetadataAtt.Attributes {
			attName := strToStructAttName(matt.DBName)
			attJson := strToStructAttJSON(matt.DBName)
			strModel = strModel + "\t" + attName + " string `yaml:\"" + attJson + "\"`\n"
		}

		appendFile(modelPath, strModel)
	}
}



func (amex *Amex) setupMiddlewareHeader() {
	middlewareDir := amex.serviceDirStructure["middleware"].(map[string]interface{})["dir"].(string)
	middlewareFile := middlewareDir + "/" + amex.serviceDirStructure["middleware"].(map[string]interface{})["file"].(string)

	str := "import (\n"
	str = str + "\t\"database/sql\"\n"
	str = str + "\t\"log\"\n\n"
	str = str + ")\n"

	appendFile(middlewareFile, str)

}
*/
