package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "felipe:admin123@tcp(localhost:3308)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
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
	// products := []Product{
	// 	{Name: "Notebook", Price: 1999.0},
	// 	{Name: "Keyboard", Price: 150.0},
	// 	{Name: "Mouse", Price: 48.0},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// fmt.Printf("Products %v\n", products)

	// where
	// var p Product
	// db.Where("name = ?", "Keyboard").First(&p)
	// fmt.Println(p)

	// update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var products []Product
	// db.Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// delete
	// var p Product
	// db.First(&p, 1)
	// db.Delete(&p)

}
