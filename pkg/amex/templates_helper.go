package amex

/*
const Insert = "// insert one {{ .modelToLower}} in the DB\n" +
	"func insert{{ .modelName}}({{ .modelToLower}} models.{{ .modelName}}) int64 {\n\n" +

	"\tvar id int64\n" +
	"\tdb := createConnection()\n" +
	"\tdefer db.Close()\n\n" +

	"\tsqlStatement := `INSERT INTO {{ .dbTableName}} ({{ .columns}}) VALUES ({{ .values}}) RETURNING ID`\n\n" +
	"\terr := db.QueryRow(sqlStatement, {{ .structValues}}).Scan(&id)\n" +

	"\tif err != nil {\n" +
	"\t\tlog.Fatalf(\"Unable to execute the query. %v\", err)\n" +
	"\t}\n\n" +

	"\treturn id\n" +
	"}\n"

const CreateConnection = "// create connection with db\n" +
	"func createConnection() *sql.DB {\n" +
	"\n\t// Open the connection\n" +
	"\tdb, err := sql.Open(\"mysql\", \"\")\n" +
	"\tif err != nil {\n" +
	"\t\tpanic(err)\n" +
	"\t}\n" +
	"\n\t// check the connection\n" +
	"\terr = db.Ping()\n" +
	"\tif err != nil {\n" +
	"\t\tpanic(err)\n" +
	"\t}\n" +
	"\n\t// return the connection\n" +
	"\treturn db\n" +
	"}\n"

func (amex *Amex) addDBConnectionFunc() {

	middlewareDir := amex.serviceDirStructure["middleware"].(map[string]interface{})["dir"].(string)
	middlewareFile := middlewareDir + "/" + amex.serviceDirStructure["middleware"].(map[string]interface{})["file"].(string)

	dbConn, err := template.New("dbConn").Parse(templates.CreateConnection)
	if err != nil {
		log.Fatal(err)
	}

	varmap := map[string]interface{}{}
	var b bytes.Buffer
	dbConn.Execute(&b, varmap)

	appendFile(middlewareFile, "\n"+b.String())
}

func (amex *Amex) addInsertFunc() {
	middlewareDir := amex.serviceDirStructure["middleware"].(map[string]interface{})["dir"].(string)
	middlewareFile := middlewareDir + "/" + amex.serviceDirStructure["middleware"].(map[string]interface{})["file"].(string)

	for _, entity := range amex.serviceData.Spec.Entities {
		modelToLower := strings.ToLower(strToStructName(entity.Name))

		dbTableName := strToDbTableName(entity.Plural)
		modelName := strToStructName(entity.Singular)

		columns := ""
		values := ""
		structValues := ""

		for i, att := range entity.Attributes {
			dbColumn := strToDbColumnName(att.Name)

			if columns == "" {
				columns = dbColumn
				values = "$" + strconv.Itoa((i + 1))
				structValues = modelToLower + "." + strToStructAttName(att.Name)
			} else {
				columns = columns + ", " + dbColumn
				values = values + ", $" + strconv.Itoa((i + 1))
				structValues = structValues + ", " + modelToLower + "." + strToStructAttName(att.Name)
			}
		}

		insert, err := template.New("insert").Parse(templates.Insert)
		if err != nil {
			log.Fatal(err)
		}

		varmap := map[string]interface{}{
			"modelToLower": modelToLower,
			"columns":      columns,
			"values":       values,
			"structValues": structValues,
			"dbTableName":  dbTableName,
			"modelName":    modelName,
		}
		var b bytes.Buffer
		insert.Execute(&b, varmap)

		appendFile(middlewareFile, "\n"+b.String())

	}
}
*/
