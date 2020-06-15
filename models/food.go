package models

type Food struct {
	ID          int `gorm:"primary_key;column:ID"`
	FoodGroupID int `gorm:"column:FoodGroupID"`
	FoodGroup   FoodGroup
	Name        string `gorm:"column:Name;"`
	NameGr      string `gorm:"column:NameGR"`
}

func (Food) TableName() string {
	return "food"
}

func GetFoods() ([]*Food, error) {
	foods := make([]*Food, 0)
	if result := VDB.Preload("FoodGroup").Limit(10).Find(&foods); result.Error != nil {
		return nil, result.Error
	}
	return foods, nil
}

func GetFoodById(id int) (*Food, error) {
	fd := &Food{}
	if result := VDB.Preload("FoodGroup").First(&fd, id); result.Error != nil {
		return nil, result.Error
	}
	return fd, nil
}

func GetFoodsByName(name string, limitCount int, offsetCount int) ([]*Food, error) {
	foods := make([]*Food, 0)
	if result := VDB.Preload("FoodGroup").Limit(limitCount).Offset(offsetCount).Where("Name LIKE ?", name+"%").Find(&foods); result.Error != nil {
		return nil, result.Error
	}
	if len(foods) == 0 {
		if result := VDB.Preload("FoodGroup").Limit(limitCount).Offset(offsetCount).Where("Name LIKE ?", "%"+name+"%").Find(&foods); result.Error != nil {
			return nil, result.Error
		}
		return foods, nil
	}
	return foods, nil
}

func GetFoodsByNameGr(nameGr string, limitCount int, offsetCount int) ([]*Food, error) {
	foods := make([]*Food, 0)
	if result := VDB.Preload("FoodGroup").Limit(limitCount).Offset(offsetCount).Where("NameGR LIKE ?", nameGr+"%").Find(&foods); result.Error != nil {
		return nil, result.Error
	}
	if len(foods) == 0 {
		if result := VDB.Preload("FoodGroup").Limit(limitCount).Offset(offsetCount).Where("NameGR LIKE ?", "%"+nameGr+"%").Find(&foods); result.Error != nil {
			return nil, result.Error
		}
		return foods, nil
	}
	return foods, nil
}
