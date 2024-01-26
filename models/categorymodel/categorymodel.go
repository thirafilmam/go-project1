package categorymodel

import (
	"go-native/config"
	"go-native/entities"
)

func GetAll() ([]entities.Category, error) {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}
