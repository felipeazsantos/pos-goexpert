package repository

import (
	"context"
	"database/sql"

	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/db"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(db *sql.DB, queries *db.Queries) CategoryRepositoryInterface {
	return &CategoryRepository{
		DB:      db,
		Queries: queries,
	}
}

func (c *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return c.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
