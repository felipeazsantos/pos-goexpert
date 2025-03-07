package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "felipe:admin123@tcp(localhost:3308)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.0,
	// })

	// create batch

	products := []Product{
		{Name: "Notebook", Price: 1999.0},
		{Name: "Keyboard", Price: 150.0},
		{Name: "Mouse", Price: 48.0},
	}

	db.Create(&products)
}
