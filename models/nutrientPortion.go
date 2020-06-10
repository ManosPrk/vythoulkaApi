package models

type NutrientPortion struct {
	ID         string `gorm:"primary_key;column:ID"`
	Group      string `gorm:"column:Group"`
	GroupGr    string `gorm:"column:GroupGR"`
	Unit       string `gorm:"column:Unit"`
	Value      string `gorm:"column:Value"`
	NutrientID int    `gorm:"column:NutrientID"`
	Nutrient   Nutrient
	FoodID     int `gorm:"column:FoodID"`
	Food       Food
}

func (NutrientPortion) TableName() string {
	return "nutrientportions"
}

func GetNutrientPortions() ([]*NutrientPortion, error) {
	nutrientPortions := make([]*NutrientPortion, 0)
	if result := VDB.Preload("Nutrient").Preload("Food.FoodGroup").Limit(10).Find(&nutrientPortions); result.Error != nil {
		return nil, result.Error
	}
	return nutrientPortions, nil
}

func GetNutrientPortionById(id string) (*NutrientPortion, error) {
	ntp := &NutrientPortion{}
	if result := VDB.Preload("Nutrient").Preload("Food.FoodGroup").First(&ntp, id); result.Error != nil {
		return nil, result.Error
	}
	return ntp, nil
}
