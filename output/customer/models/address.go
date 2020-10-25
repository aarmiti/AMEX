package models

type Address struct {
	Id string `yaml:"id"`

	//Main attributes
	PostCode int `yaml:"postCode"`
	State string `yaml:"state"`
	Street string `yaml:"street"`
	Number int `yaml:"number"`
	IsDefault bool `yaml:"isDefault"`

	//Metadata attributes
	CreatedDate string `yaml:"createdDate"`
	CreatedBy string `yaml:"createdBy"`
	CreatedSource string `yaml:"createdSource"`
	LastModifiedDate string `yaml:"lastModifiedDate"`
	LastModifiedBy string `yaml:"lastModifiedBy"`
	LastModifiedSource string `yaml:"lastModifiedSource"`
	Status string `yaml:"status"`
	Version string `yaml:"version"`

	//Relationship attributes
	CustomerEList []Customer `json:"customerEList"`
}