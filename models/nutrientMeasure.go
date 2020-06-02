package models

type NutrientMeasure struct {
	ID         string           `json:"id"`
	Label      string           `json:"label"`
	LabelGr    string           `json:"labelGr"`
	Quantity   string           `json:"quantity"`
	Value      string           `json:"value"`
	EquivUnit  string           `json:"equivUnit"`
	Equivalent string           `json:"equivalent"`
	Portion    *NutrientPortion `json:"nutrientPortion"`
}
