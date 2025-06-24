package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/db"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer dbt.Close()

	dbt.Exec("CREATE TABLE category (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT)")
	dbt.Exec("CREATE TABLE course (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT, category_id INTEGER)")

	categoryRepository := repository.NewCategoryRepository(dbt, db.New(dbt))
	courseRepository := repository.NewCourseRepository(dbt, db.New(dbt))

	useCase := NewAddCourseUseCase(categoryRepository, courseRepository)

	input := InputUseCase{
		CategoryName:     "Go",
		CourseName:       "Go Expert",
		CourseCategoryID: 1,
	}

	err = useCase.Execute(context.Background(), input)
	assert.NoError(t, err)

	var category db.Category
	dbt.QueryRow("SELECT * FROM category WHERE name = ?", input.CategoryName).Scan(&category.ID, &category.Name, &category.Description)
	assert.Equal(t, input.CategoryName, category.Name)
	assert.Equal(t, "", category.Description.String)
	assert.Equal(t, input.CourseCategoryID, int(category.ID))

	var course db.Course
	dbt.QueryRow("SELECT * FROM course WHERE name = ?", input.CourseName).Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
	assert.Equal(t, input.CourseName, course.Name)
	assert.Equal(t, "", course.Description.String)
	assert.Equal(t, input.CourseCategoryID, int(course.CategoryID.Int32))
}
