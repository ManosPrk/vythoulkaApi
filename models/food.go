package models

import "fmt"

type Food struct {
	ID     string     `json:"id"`
	ApiId  string     `json:"apiId"`
	Name   string     `json:"name"`
	NameGR string     `json:"nameGr"`
	Group  *FoodGroup `json:"foodGroup"`
}

func newFood() *Food {
	return &Food{
		Group: new(FoodGroup),
	}
}

func GetFoods() ([]*Food, error) {
	rows, err := VDB.Query("SELECT food.ID, food.ApiId, food.Name, food.NameGR, foodgroups.*  FROM food INNER JOIN foodgroups ON food.FoodGroupID=foodgroups.ID LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := make([]*Food, 0)
	for rows.Next() {
		fd := newFood()
		err := rows.Scan(&fd.ID, &fd.ApiId, &fd.Name, &fd.NameGR, &fd.Group.ID, &fd.Group.ApiId, &fd.Group.Name)
		if err != nil {
			return nil, err
		}
		foods = append(foods, fd)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return foods, nil
}

func GetFoodById(id string) (*Food, error) {
	fd := newFood()
	err := VDB.QueryRow("SELECT food.ApiId, food.Name, food.NameGR, foodgroups.Name, foodgroups.ApiId  FROM food INNER JOIN foodgroups ON food.FoodGroupID=foodgroups.ID WHERE food.ApiId=?", id).Scan(&fd.ID, &fd.Name, &fd.NameGR, &fd.Group.Name, &fd.Group.ID)
	if err != nil {
		return nil, err
	}
	return fd, nil
}

func GetFoodsByName(name string) ([]*Food, error) {
	query := fmt.Sprint("SELECT food.ApiId, food.Name, food.NameGR, foodgroups.Name, foodgroups.ApiId  FROM food INNER JOIN foodgroups ON food.FoodGroupID=foodgroups.ID WHERE food.Name LIKE ? LIMIT 10")
	rows, err := VDB.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := make([]*Food, 0)
	for rows.Next() {
		fd := newFood()
		err := rows.Scan(&fd.ID, &fd.Name, &fd.NameGR, &fd.Group.Name, &fd.Group.ID)
		if err != nil {
			return nil, err
		}
		foods = append(foods, fd)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return foods, nil
}

func GetFoodsByNameGr(nameGr string) ([]*Food, error) {
	query := fmt.Sprint("SELECT food.ApiId, food.Name, food.NameGR, foodgroups.Name, foodgroups.ApiId  FROM food INNER JOIN foodgroups ON food.FoodGroupID=foodgroups.ID WHERE food.NameGr LIKE ? LIMIT 10")
	rows, err := VDB.Query(query, "%"+nameGr+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := make([]*Food, 0)
	for rows.Next() {
		fd := newFood()
		err := rows.Scan(&fd.ID, &fd.Name, &fd.NameGR, &fd.Group.Name, &fd.Group.ID)
		if err != nil {
			return nil, err
		}
		foods = append(foods, fd)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return foods, nil
}
