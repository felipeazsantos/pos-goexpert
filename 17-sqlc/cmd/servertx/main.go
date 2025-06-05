package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/felipeazsantos/pos-goexpert/17-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	db *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		db:      dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error rolling back transaction: %v, original error: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) createCourseAndCategory(ctx context.Context, categoryParams db.CreateCategoryParams, courseParams db.CreateCourseParams) error {
	return c.callTx(ctx, func(q *db.Queries) error {
		if err := q.CreateCategory(ctx, categoryParams); err != nil {
			return fmt.Errorf("error creating category: %w", err)
		}

		if err := q.CreateCourse(ctx, courseParams); err != nil {
			return fmt.Errorf("error creating course: %w", err)
		}

		return nil
	})
}

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	courseDB := NewCourseDB(conn)
	err = courseDB.createCourseAndCategory(ctx, db.CreateCategoryParams{
		Name: "Go",
		Description: sql.NullString{String: "Go programming language", Valid: true},
	}, db.CreateCourseParams{
		Name:        "Go Expert",
		Description: sql.NullString{String: "Become an expert in Go", Valid: true},
		CategoryID:  sql.NullInt32{Int32: 2, Valid: true},
	})
	if err != nil {
		panic(err)
	}

}
