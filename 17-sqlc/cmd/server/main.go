package main

import (
	"context"
	"database/sql"

	"github.com/felipeazsantos/pos-goexpert/17-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: "Category 1",
		Description: sql.NullString{
			String: "Description of Category 1",
			Valid:  true,
		},
	})

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}