package repository

import (
	"github.com/akbarpambudi/shoppies/product"
	"github.com/akbarpambudi/shoppies/product/errors"
)

type InMemoryProductRepository struct {
	products map[string]*product.Product
}

func (i *InMemoryProductRepository) FindById(id string) (*product.Product, error) {
	product := i.products[id]
	if product == nil {
		return nil, errors.ProductNotFoundErr
	}
	return product, nil
}

func (i *InMemoryProductRepository) Add(product product.Product) error {
	persistedProduct := i.products[product.Id]
	if persistedProduct != nil {
		return errors.ProductAlreadyExistsErr
	}
	i.products[product.Id] = &product
	return nil
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{products: map[string]*product.Product{}}
}
