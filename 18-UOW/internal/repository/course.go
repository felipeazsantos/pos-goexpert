package repository

import (
	"context"
	"database/sql"

	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/db"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(db *sql.DB, queries *db.Queries) CourseRepositoryInterface {
	return &CourseRepository{
		DB:      db,
		Queries: queries,
	}
}

func (c *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return c.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: sql.NullInt32{Int32: int32(course.CategoryID), Valid: true},
	})
}
