package models

type Nutrient struct {
	ID     string
	ApiId  string `json:"id"`
	Name   string `json:"name"`
	NameGR string `json:"nameGr"`
}
