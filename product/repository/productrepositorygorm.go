package repository

import (
	"github.com/akbarpambudi/shoppies/product"
	"github.com/akbarpambudi/shoppies/product/errors"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryGorm struct {
	db *gorm.DB
}

func (repo ProductRepositoryGorm) FindById(id string) (*product.Product, error) {
	product := new(product.Product)
	if err := repo.db.Where("id =?", id).Find(product).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.ProductNotFoundErr
		}
		return nil, err
	}

	return product, nil
}

func (repo ProductRepositoryGorm) Add(product product.Product) error {
	if err := repo.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func NewProductRepositoryGorm(db *gorm.DB) *ProductRepositoryGorm {
	return &ProductRepositoryGorm{db: db}
}
