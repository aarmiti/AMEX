package amex

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func appendFile(filePath, data string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

func createDir(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func createFile(filePath string) {
	os.Remove(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func lcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

/*
# DB:
#  name   = XXX_YYY
#  column = ZZZ_WWW
#  file = ffff.sql
#
# Struct:
#  name = XxxYyy
#  att  = IiiPppp
#  json = rrrTtttHhhh
#
# Go file:
# name = ffff.go
*/
func strToDbFileName(str string) string {
	s := strings.Split(str, "_")
	f := ""
	for i := 0; i < len(s); i++ {
		f += strings.ToLower(s[i])
	}
	return f + ".sql"
}
func strToStructAttName(str string) string {
	return strToStructName(str)
}
func strToStructAttJSON(str string) string {
	str = strings.ToLower(str)
	s := strings.Split(str, "_")
	f := ""
	for i := 0; i < len(s); i++ {
		f += strings.Title(s[i])
	}
	return lcFirst(f)
}

func strToDbTableName(str string) string {
	return strings.ToUpper(str)
}
func strToDbColumnName(str string) string {
	return strings.ToUpper(str)
}

func strToGoFileName(str string) string {
	s := strings.Split(str, "_")
	f := ""
	for i := 0; i < len(s); i++ {
		f += strings.ToLower(s[i])
	}
	return f + ".go"
}

func strToStructName(str string) string {
	str = strings.ToLower(str)
	s := strings.Split(str, "_")
	f := ""
	for i := 0; i < len(s); i++ {
		f += strings.Title(s[i])
	}
	return f
}
