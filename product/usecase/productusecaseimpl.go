package usecase

import (
	"math/big"

	"github.com/akbarpambudi/shoppies/product"
	"github.com/google/uuid"
)

type ProductUseCaseImpl struct {
	repository product.ProductRepository
}

func (p *ProductUseCaseImpl) FindProductById(id string) (*product.Product, error) {
	return p.repository.FindById(id)
}

func (p *ProductUseCaseImpl) AddProduct(productName string, stock int, price int) error {
	return p.repository.Add(product.Product{Id: uuid.New().String(), Name: productName, InStock: stock, Price: big.NewInt(int64(price))})
}

func NewProductUseCaseImpl(repository product.ProductRepository) *ProductUseCaseImpl {
	return &ProductUseCaseImpl{repository: repository}
}
