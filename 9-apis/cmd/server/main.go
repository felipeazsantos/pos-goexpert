package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/felipeazsantos/pos-goexpert/apis/configs"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/dto"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})


	productDB := database.NewProduct(db)
	productHandler := NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct, err := entity.NewProduct(product.Name, product.Description, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}
