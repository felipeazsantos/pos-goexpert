package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "felipe:admin123@tcp(localhost:3308)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	gormRelations(db)

}

func gormRelations(db *gorm.DB) {
	// create category
	// category := Category{Name: "Eletronics"}
	// db.Create(&category)

	// create prdocut
	product := Product{
		Name:       "Mouse",
		Price:      1000.0,
		CategoryID: 1,
		// CategoryID: category.ID,
	}
	db.Create(&product)

	// create a serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category
	err := db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Println(product.Name, category.Name)
		}
	}
}

func gormCrud(db *gorm.DB) {
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.0,
	})

	// create batch
	productsCreateBatch := []Product{
		{Name: "Notebook", Price: 1999.0},
		{Name: "Keyboard", Price: 150.0},
		{Name: "Mouse", Price: 48.0},
	}
	db.Create(&productsCreateBatch)

	// select one
	var productSelecOne Product
	db.First(&productSelecOne, 2)
	db.First(&productSelecOne, "name = ?", "Mouse")
	fmt.Println(productSelecOne)

	// select all
	var productsSelectAll []Product
	db.Find(&productsSelectAll)
	fmt.Printf("Products %v\n", productsSelectAll)

	// where
	var productWhere Product
	db.Where("name = ?", "Keyboard").First(&productWhere)
	fmt.Println(productWhere)

	// update
	var productUpdate Product
	db.First(&productUpdate, 1)
	productUpdate.Name = "New Mouse"
	db.Save(&productUpdate)

	var productsUpdate []Product
	db.Find(&productsUpdate)

	for _, product := range productsUpdate {
		fmt.Println(product)
	}

	// delete
	var productDelete Product
	db.First(&productDelete, 1)
	db.Delete(&productDelete)
}
