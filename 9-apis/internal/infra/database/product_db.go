package database

import (
	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) ProductInterface {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		err := p.DB.Limit(limit).Offset(offset).Order("created_at " + sort).Find(&products).Error
		return products, err
	}

	err := p.DB.Order("created_at " + sort).Find(&products).Error
	return products, err
}

func (p *Product) Update(product *entity.Product) error {
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}
