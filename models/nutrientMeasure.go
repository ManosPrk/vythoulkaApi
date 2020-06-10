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

// func newNutrientMeasure() *NutrientMeasure {
// 	return &NutrientMeasure{
// 		Portion: NewNutrientPortion(),
// 	}
// }

// func GetNutrientMeasures() ([]*NutrientMeasure, error) {
// 	rows, err := VDB.Query("SELECT nutrientMeasures.*, nutrientportions.*, food.*, nutrients.* FROM nutrientmeasures INNER JOIN nutrientportions ON nutrientmeasures.NutrientPortionID=nutrientportions.ID INNER JOIN food ON nutrientportions.FoodID=food.ID INNER JOIN foodgroups ON food.FoodGroupID=foodgroups.ID INNER JOIN nutrients ON nutrientportions.NutrientID=nutrients.ID LIMIT 10")

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	nutrientMeasures := make([]*NutrientMeasure, 0)
// 	for rows.Next() {
// 		ntm := newNutrientMeasure()
// 		err := rows.Scan(&ntm.ID, &ntm.Label, &ntm.LabelGr, &ntm.Quantity, &ntm.Value, &ntm.EquivUnit, &ntm.EquivUnit, &ntm.Equivalent, &ntm.Portion.ID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		nutrientMeasures = append(nutrientMeasures, ntm)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return nutrientMeasures, nil
// }

// func GetNutrientMeasureById(id string) (*NutrientMeasure, error) {
// 	nt := newNutrientMeasure()
// 	err := VDB.QueryRow("SELECT nutrientMeasures.ID, nutrientMeasures.Name, nutrientMeasures.NameGR FROM nutrientMeasures WHERE nutrientMeasures.ID=?", id).Scan(&nt.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nt, nil
// }

// func GetNutrientMeasuresByName(name string) ([]*NutrientMeasure, error) {
// 	query := fmt.Sprint("SELECT nutrientMeasures.ID, nutrientMeasures.Name, nutrientMeasures.NameGR FROM nutrientMeasures WHERE nutrientMeasures.Name LIKE ? LIMIT 10")
// 	rows, err := VDB.Query(query, "%"+name+"%")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	nutrientMeasures := make([]*NutrientMeasure, 0)
// 	for rows.Next() {
// 		nt := newNutrientMeasure()
// 		err := rows.Scan(&nt.ID, &nt.Name, &nt.NameGR)
// 		if err != nil {
// 			return nil, err
// 		}
// 		nutrientMeasures = append(nutrientMeasures, nt)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return nutrientMeasures, nil
// }

// func GetNutrientMeasuresByNameGr(nameGr string) ([]*NutrientMeasure, error) {
// 	query := fmt.Sprint("SELECT nutrientMeasures.ID, nutrientMeasures.Name, nutrientMeasures.NameGR FROM nutrientMeasures WHERE nutrientMeasures.NameGr LIKE ? LIMIT 10")
// 	rows, err := VDB.Query(query, "%"+nameGr+"%")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	nutrientMeasures := make([]*NutrientMeasure, 0)
// 	for rows.Next() {
// 		nt := newNutrientMeasure()
// 		err := rows.Scan(&nt.ID, &nt.Name, &nt.NameGR)
// 		if err != nil {
// 			return nil, err
// 		}
// 		nutrientMeasures = append(nutrientMeasures, nt)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return nutrientMeasures, nil
// }
