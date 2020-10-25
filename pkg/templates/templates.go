package templates

const Insert = "// insert one {{ .modelToLower}} in the DB\n" +
	"func insert{{ .modelTitle}}({{ .modelToLower}} models.{{ .modelTitle}}) int64 {\n\n" +
	"\t// create the mysql db connection\n" +
	"\tdb := createConnection()\n\n" +
	"\t// close the db connection\n" +
	"\tdefer db.Close()\n\n" +
	"\t// create the insert sql query\n" +
	"\t// returning userid will return the id of the inserted user\n" +
	"\tsqlStatement := `INSERT INTO {{ .modelToLower}} ({{ .columns}}) VALUES ({{ .values}}) RETURNING {{ .Name}}`\n\n" +
	"\t// the inserted id will store in this id\n" +
	"\tvar id int64\n\n" +
	"\t// execute the sql statement\n" +
	"\t// Scan function will save the insert id in the id\n" +
	"\terr := db.QueryRow(sqlStatement, {{ .structValues}}).Scan(&id)\n" +
	"\tif err != nil {\n" +
	"\t\tlog.Fatalf(\"Unable to execute the query. %v\", err)\n" +
	"\t}\n\n" +
	"\tfmt.Printf(\"Inserted a single record %v\", id)\n\n" +
	"\t// return the inserted id\n" +
	"\treturn id\n" +
	"}"
