package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(name, description, categoryID string) (Course, error) {
	id := uuid.New().String()
	query := "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)"

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return Course{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) GetAll() ([]Course, error) {
	query := "SELECT id, name, description, category_id FROM courses"
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	query := "SELECT id, name, description, category_id FROM courses WHERE category_id = $1"
	rows, err := c.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}