package models

type City struct {

	//Persistent attributes
	Name string `json:"name"`

	//Relational attributes
	Detail Detail `json:"detail"`
}