package models

type Tag struct {

	//Persistent attributes
	Name string `json:"name"`

	//Relational attributes
	Tasks []Task `json:"tasks"`
}