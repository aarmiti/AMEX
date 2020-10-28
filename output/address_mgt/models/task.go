package models

type Task struct {

	//Persistent attributes
	Name string `json:"name"`

	//Relational attributes
	Tags []Tag `json:"tags"`
}