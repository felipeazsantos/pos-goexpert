package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/db"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/repository"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/pkg/uow"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer dbt.Close()

	dbt.Exec("CREATE TABLE category (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT)")
	dbt.Exec("CREATE TABLE course (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT, category_id INTEGER)")

	uow := uow.NewUow(dbt)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		return repository.NewCategoryRepository(dbt, db.New(tx))
	})
	uow.Register("CourseRepository", func (tx *sql.Tx) interface{}  {
		return repository.NewCourseRepository(dbt, db.New(tx))
	})

	useCase := NewAddCourseUseCaseUow(uow)

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
