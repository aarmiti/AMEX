package models

type Customer struct {
	Id string `yaml:"id"`

	//Main attributes
	FirstName string `yaml:"firstName"`
	LastName string `yaml:"lastName"`
	Email string `yaml:"email"`

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
	ParentEObj Customer `json:"parentEObj"`
	DetailEObj Detail `json:"detailEObj"`
	AddressEList []Address `json:"addressEList"`
}