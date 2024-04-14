package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{DB: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.DB.Query("SELECT id, name, description FROM categories")

	if err != nil {
		return []Category{}, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)

		if err != nil {
			return []Category{}, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
