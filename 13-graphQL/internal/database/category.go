package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	query := "INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)"

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return Category{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) GetAll() ([]Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	query := `
		SELECT c.id, c.name, c.description
		FROM categories c
		JOIN courses co ON c.id = co.category_id
		WHERE co.id = $1`
	rows, err := c.db.Query(query, courseID)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	var category Category
	if rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return Category{}, err
		}
	} else {
		return Category{}, sql.ErrNoRows
	}

	if err := rows.Err(); err != nil {
		return Category{}, err
	}
	return category, nil
}