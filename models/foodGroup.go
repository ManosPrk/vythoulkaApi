package models

type FoodGroup struct {
	ID     string
	ApiId  string `json:"apiId"`
	Name   string `json:"name"`
	NameGR string `json:"nameGr"`
}
