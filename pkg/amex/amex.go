package amex

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

// MetaData Attributes
type MetaDataAttributes struct {
	Attributes []MetaAttributes
}
type MetaAttributes struct {
	DBName     string
	StructName string
	DataType   string
	Defualt    string
	Nullable   string
}

var MetaDataAttributesList = &MetaDataAttributes{
	Attributes: []MetaAttributes{
		MetaAttributes{
			DBName:   "CREATED_DATE",
			DataType: "TIMESTAMP",
			Defualt:  "CURRENT_TIMESTAMP",
			Nullable: "",
		},
		MetaAttributes{
			DBName:   "CREATED_BY",
			DataType: "VARCHAR(100)",
			Defualt:  "",
			Nullable: "",
		},
		MetaAttributes{
			DBName:   "CREATED_SOURCE",
			DataType: "VARCHAR(100)",
			Defualt:  "",
			Nullable: "",
		},
		MetaAttributes{
			DBName:   "LAST_MODIFIED_DATE",
			DataType: "TIMESTAMP",
			Defualt:  "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			Nullable: "NOT NULL",
		},
		MetaAttributes{
			DBName:   "LAST_MODIFIED_BY",
			DataType: "VARCHAR(100)",
			Defualt:  "",
			Nullable: "",
		},
		MetaAttributes{
			DBName:   "LAST_MODIFIED_SOURCE",
			DataType: "VARCHAR(100)",
			Defualt:  "",
			Nullable: "",
		},
		MetaAttributes{
			DBName:   "STATUS",
			DataType: "VARCHAR(50)",
			Defualt:  "",
			Nullable: "",
		},
		MetaAttributes{
			DBName:   "VERSION",
			DataType: "INT",
			Defualt:  "",
			Nullable: "",
		},
	},
}

// MicroService Manifest
type MicroService struct {
	Kind     string   `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
	Spec     Spec     `yaml:"spec"`
}

type Metadata struct {
	Name         string `yaml:"name"`
	StorageSpace string `yaml:"storageSpace"`
}

type Spec struct {
	Objects []Object `yaml:"objects"`
}

type Object struct {
	Name       string     `yaml:"name"`
	Singular   string     `yaml:"singular"`
	Plural     string     `yaml:"plural"`
	Attributes Attributes `yaml:"attributes"`
}

type Attributes struct {
	Persistent []Persistent `yaml:"persistent"`
	Relational []Relational `yaml:"relational"`
	Calculated []Calculated `yaml:"calculated"`
	External   []External   `yaml:"external"`
}

type Persistent struct {
	Name           string         `yaml:"name"`
	PersistentSpec PersistentSpec `yaml:"spec"`
}

type PersistentSpec struct {
	DataType string `yaml:"dataType"`
	Required bool   `yaml:"required"`
	Length   string `yaml:"length"`
	Default  string `yaml:"default"`
}

type Relational struct {
	Name   string `yaml:"name"`
	Object string `yaml:"object"`
	Type   string `yaml:"type"`
}
type Calculated struct {
	Name string `yaml:"name"`
}
type External struct {
	Name string `yaml:"name"`
}

type AmexEngine struct {
	microServiceData        *MicroService
	microServicDirStructure map[string]interface{}
}

func NewAmexEngine(path, outputPath string) *AmexEngine {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	m := &MicroService{}
	err = yaml.Unmarshal(buf, m)
	if err != nil {
		log.Fatal(fmt.Errorf("in file %q: %w", path, err))

		return nil
	}

	microservicePath := outputPath + "/" + strings.ToLower(m.Metadata.Name)
	return &AmexEngine{
		microServiceData: m,
		microServicDirStructure: map[string]interface{}{
			"main": map[string]interface{}{
				"dir":  microservicePath,
				"name": "main",
				"file": "main.go",
			},
			"middleware": map[string]interface{}{
				"dir":  microservicePath + "/middleware",
				"name": "middleware",
				"file": "handlers.go",
			},
			"models": map[string]interface{}{
				"dir":  microservicePath + "/models",
				"name": "models",
				"file": "",
			},
			"router": map[string]interface{}{
				"dir":  microservicePath + "/router",
				"name": "router",
				"file": "router.go",
			},
			"schema": map[string]interface{}{
				"dir":  microservicePath + "/schema",
				"name": "",
				"file": "",
			},
		},
	}
}

func (amexEngine *AmexEngine) SetupMicroService() {

	// Create Micro Service File Structure
	amexEngine.createMicroServiceFileStructure()

	// Create Model Files
	amexEngine.createModelFiles()

	// Create Storage Space Files
	amexEngine.createStorageSpaceFiles()

	// 	/* #################################################################### */
	// 	// Add middleware headers
	// 	// Add db connection func
	// 	amex.setupMiddlewareHeader()
	// 	amex.addDBConnectionFunc()
	// 	amex.addInsertFunc()
}
