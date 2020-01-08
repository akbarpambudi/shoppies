package product

import (
	"math/big"
)

type Product struct {
	Id      string `gorm:"primary_key"`
	Name    string
	Price   *big.Int
	InStock int
}

func (p Product) TableName() string {
	return "s_product"
}
