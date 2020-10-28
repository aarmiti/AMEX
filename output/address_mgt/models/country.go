package models

type Country struct {

	//Persistent attributes
	CName string `json:"cName"`

	//Relational attributes
	Cities []City `json:"cities"`
}