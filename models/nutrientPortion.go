package models

type NutrientPortion struct {
	ID       string    `json:"id"`
	Group    string    `json:"Group"`
	GroupGr  string    `json:"GroupGr"`
	Unit     string    `json:"unit"`
	Value    string    `json:"value"`
	Nutrient *Nutrient `json:"nutrient"`
	Food     *Food     `json:"food"`
}
