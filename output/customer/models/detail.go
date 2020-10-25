package models

type Detail struct {
	Id string `yaml:"id"`

	//Main attributes
	DetailName string `yaml:"detailName"`

	//Metadata attributes
	CreatedDate string `yaml:"createdDate"`
	CreatedBy string `yaml:"createdBy"`
	CreatedSource string `yaml:"createdSource"`
	LastModifiedDate string `yaml:"lastModifiedDate"`
	LastModifiedBy string `yaml:"lastModifiedBy"`
	LastModifiedSource string `yaml:"lastModifiedSource"`
	Status string `yaml:"status"`
	Version string `yaml:"version"`
}