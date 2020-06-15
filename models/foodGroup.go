package models

type FoodGroup struct {
	ID     int    `gorm:"primary_key;column:ID"`
	Name   string `gorm:"column:Name"`
	NameGr string `gorm:"column:NameGR"`
}

func (FoodGroup) TableName() string {
	return "foodgroups"
}
