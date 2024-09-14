package categorymodel

import (
	"CRUD-Project1/config"
	"CRUD-Project1/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`select * from categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var Category entities.Category
		rows.Scan(&Category.Id, &Category.Name, &Category.CreatedAt, &Category.UpdatedAt)

		if err != nil {
			panic(err)
		}

		categories = append(categories, Category)
	}
	return categories

}

func Create(Category entities.Category) bool {

	result, err := config.DB.Exec(`
		INSERT INTO categories (name, created_at, updated_at)
		VALUE (?,?,?)`,
		Category.Name, Category.CreatedAt, Category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return LastInsertId > 0

}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`
	SELECT id, name FROM categories WHERE id = ?
	`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}
	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`
	UPDATE categories SET name = ?, Updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
