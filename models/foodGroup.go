package models

type FoodGroup struct {
	ID     int    `gorm:"primary_key;column:ID"`
	Name   string `gorm:"column:Name"`
	NameGr string `gorm:"column:NameGR"`
}

func (FoodGroup) TableName() string {
	return "foodgroups"
}

func newFoodGroup() *FoodGroup {
	return &FoodGroup{}
}

func GetFoodGroups() ([]*FoodGroup, error) {
	fdg := make([]*FoodGroup, 0)
	VDB.Limit(10).Find(&fdg)
	return fdg, nil
}
