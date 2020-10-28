package models

type Detail struct {

	//Persistent attributes
	Name string `json:"name"`

	//Relational attributes
	Parent Detail `json:"parent"`
}