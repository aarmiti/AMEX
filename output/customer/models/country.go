package models

type Country struct {
	Id string `yaml:"id"`

	//Main attributes
	Name string `yaml:"name"`

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
	CityEList []City `json:"cityEList"`
}