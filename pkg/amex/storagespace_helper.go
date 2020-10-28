package amex

import (
	"strings"
)

func (amex *AmexEngine) createStorageSpaceFiles() {
	storageSpace := amex.microServiceData.Metadata.StorageSpace
	if storageSpace == "mysql" {
		amex.createMySqlFiles()
	}
}

func (amex *AmexEngine) createMySqlFiles() {
	schemaDir := amex.microServicDirStructure["schema"].(map[string]interface{})["dir"].(string)
	dbSchemaName := strings.ToUpper(amex.microServiceData.Metadata.Name)

	dbFilesMap := make(map[string]string)
	dbTablesMap := make(map[string]string)

	// Persistent columns
	for _, object := range amex.microServiceData.Spec.Objects {
		if len(object.Attributes.Persistent) > 0 || len(object.Attributes.Relational) > 0 || len(object.Attributes.External) > 0 {
			dbFileName := strToDbFileName(object.Plural)
			dbTableName := strToDbTableName(object.Plural)

			dbFilesMap[object.Name] = dbFileName
			dbTablesMap[object.Name] = dbTableName

			dbFilePath := schemaDir + "/" + dbFileName
			createFile(dbFilePath)

			strDB := "CREATE SCHEMA IF NOT EXISTS `" + dbSchemaName + "` DEFAULT CHARACTER SET utf8;\n\n"
			strDB = strDB + "USE `" + dbSchemaName + "`;\n\n"

			strDB = strDB + "DROP TABLE IF EXISTS `" + dbSchemaName + "`.`" + dbTableName + "`;\n\n"
			strDB = strDB + "CREATE TABLE `" + dbSchemaName + "`.`" + dbTableName + "` (\n"
			strDB = strDB + "  `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,\n"

			strDB = strDB + "\n  -- Persistent columns --\n"

			for _, att := range object.Attributes.Persistent {
				dbColumn := strToDbColumnName(att.Name)
				dbColumn = "  `" + dbColumn + "` " + getMySqlColumnSpec(att.PersistentSpec) + ",\n"
				strDB = strDB + dbColumn
			}

			appendFile(dbFilePath, strDB)
		}
	}

	// Relational columns not N:N
	for _, val := range dbFilesMap {
		dbFilePath := schemaDir + "/" + val
		strDB := "\n  -- Relational columns --\n"
		appendFile(dbFilePath, strDB)

	}
	for _, object := range amex.microServiceData.Spec.Objects {
		for _, att := range object.Attributes.Relational {

			if att.Type == "1:1" || att.Type == "1:N" {
				dbFilePath := schemaDir + "/" + dbFilesMap[att.Object]
				strDB := ""
				if object.Name == att.Object {
					strDB = "  `" + strToDbColumnName(att.Name+"_ID") + "` BIGINT UNSIGNED,\n"
				} else {
					strDB = "  `" + strToDbColumnName(object.Singular+"_ID") + "` BIGINT UNSIGNED,\n"
				}
				appendFile(dbFilePath, strDB)
			}
		}
	}

	// Constraint
	for _, val := range dbFilesMap {
		dbFilePath := schemaDir + "/" + val
		strDB := "\n  -- Constraint --\n"
		appendFile(dbFilePath, strDB)
	}
	// FK
	for _, object := range amex.microServiceData.Spec.Objects {
		for _, att := range object.Attributes.Relational {

			if att.Type == "1:1" || att.Type == "1:N" {
				dbFilePath := schemaDir + "/" + dbFilesMap[att.Object]
				strDB := ""
				if object.Name == att.Object {
					strDB = "  CONSTRAINT `FK_" + strToDbColumnName(att.Name+"_ID") + "` FOREIGN KEY (`" + strToDbColumnName(att.Name+"_ID") + "`) REFERENCES `" + dbSchemaName + "`.`" + dbTablesMap[object.Name] + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"
				} else {
					strDB = "  CONSTRAINT `FK_" + strToDbColumnName(object.Singular+"_ID") + "` FOREIGN KEY (`" + strToDbColumnName(object.Singular+"_ID") + "`) REFERENCES `" + dbSchemaName + "`.`" + dbTablesMap[object.Name] + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"
				}
				appendFile(dbFilePath, strDB)
			}
		}
	}
	// PK and clode table
	for _, val := range dbFilesMap {
		dbFilePath := schemaDir + "/" + val
		strDB := "  PRIMARY KEY (`ID`) USING BTREE"
		strDB = strDB + "\n) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;\n"
		appendFile(dbFilePath, strDB)
	}

	// N:N Tables
	joinTableList := []string{}
	for _, object := range amex.microServiceData.Spec.Objects {
		for _, att := range object.Attributes.Relational {
			if att.Type == "N:N" {
				exists := false
				for _, str := range joinTableList {

					if str == (object.Name + "#" + att.Object) {
						exists = true
					}
					if str == (att.Object + "#" + object.Name) {
						exists = true
					}
				}
				if !exists {
					joinTableList = append(joinTableList, object.Name+"#"+att.Object)

					dbFileName := strToDbFileName("junction_" + dbTablesMap[object.Name] + "_with_" + dbTablesMap[att.Object])
					dbTableName := strToDbTableName("junction_" + dbTablesMap[object.Name] + "_with_" + dbTablesMap[att.Object])

					dbFilePath := schemaDir + "/" + dbFileName
					createFile(dbFilePath)

					strDB := "CREATE SCHEMA IF NOT EXISTS `" + dbSchemaName + "` DEFAULT CHARACTER SET utf8;\n\n"
					strDB = strDB + "USE `" + dbSchemaName + "`;\n\n"

					strDB = strDB + "DROP TABLE IF EXISTS `" + dbSchemaName + "`.`" + dbTableName + "`;\n\n"
					strDB = strDB + "CREATE TABLE `" + dbSchemaName + "`.`" + dbTableName + "` (\n"
					strDB = strDB + "  `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,\n"

					strDB = strDB + "\n  -- Relationship columns --\n"
					columnName := strToDbColumnName(object.Name + "_ID")
					strDB = strDB + "  `" + columnName + "` BIGINT UNSIGNED,\n"

					columnName = strToDbColumnName(att.Object + "_ID")
					strDB = strDB + "  `" + columnName + "` BIGINT UNSIGNED,\n"

					strDB = strDB + "\n  -- Constraint --\n"

					dbTablenameSingular := strToDbTableName(object.Name + "_ID")

					tablename := strToDbTableName(dbTablesMap[object.Name])
					strDB = strDB + "  CONSTRAINT `FK_" + dbTablenameSingular + "` FOREIGN KEY (`" + dbTablenameSingular + "`) REFERENCES `" + dbSchemaName + "`.`" + tablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

					dbTablenameSingular = strToDbTableName(att.Object + "_ID")

					tablename = strToDbTableName(dbTablesMap[att.Object])
					strDB = strDB + "  CONSTRAINT `FK_" + dbTablenameSingular + "` FOREIGN KEY (`" + dbTablenameSingular + "`) REFERENCES `" + dbSchemaName + "`.`" + tablename + "` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,\n"

					strDB = strDB + "  PRIMARY KEY (`ID`) USING BTREE"
					strDB = strDB + "\n) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;\n"
					appendFile(dbFilePath, strDB)

				}
			}
		}
	}

}

func getMySqlColumnSpec(att PersistentSpec) string {
	str := ""

	switch att.DataType {
	case "string":
		str = "VARCHAR"
		if att.Length != "0" {
			str = str + "(" + att.Length + ")"
		} else {
			str = str + "(255)"
		}

	case "int":
		str = "INT"

	case "bool":
		str = "BOOLEAN"
	}

	if att.Required {
		str = str + " NOT NULL"
	}
	return str
}
