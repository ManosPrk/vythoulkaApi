package models

type Nutrient struct {
	ID     int    `gorm:"primary_key;column:ID"`
	Name   string `gorm:"column:Name"`
	NameGr string `gorm:"column:NameGR"`
}

func GetNutrients() ([]*Nutrient, error) {
	nutrients := make([]*Nutrient, 0)
	if result := VDB.Limit(10).Find(&nutrients); result.Error != nil {
		return nil, result.Error
	}
	return nutrients, nil
}

func GetNutrientById(id string) (*Nutrient, error) {
	nt := &Nutrient{}
	if result := VDB.First(&nt, id); result.Error != nil {
		return nil, result.Error
	}
	return nt, nil
}

func GetNutrientsByName(name string) ([]*Nutrient, error) {
	nutrients := make([]*Nutrient, 0)
	if result := VDB.Limit(10).Where("Name LIKE ?", name+"%").Find(&nutrients); result.Error != nil {
		return nil, result.Error
	}
	if len(nutrients) == 0 {
		if result := VDB.Limit(10).Where("Name LIKE ?", "%"+name+"%").Find(&nutrients); result.Error != nil {
			return nil, result.Error
		}
		return nutrients, nil
	}
	return nutrients, nil
}

func GetNutrientsByNameGr(nameGr string) ([]*Nutrient, error) {
	nutrients := make([]*Nutrient, 0)
	if result := VDB.Limit(10).Where("NameGR LIKE ?", nameGr+"%").Find(&nutrients); result.Error != nil {
		return nil, result.Error
	}
	if len(nutrients) == 0 {
		if result := VDB.Limit(10).Where("NameGR LIKE ?", "%"+nameGr+"%").Find(&nutrients); result.Error != nil {
			return nil, result.Error
		}
		return nutrients, nil
	}
	return nutrients, nil
}
