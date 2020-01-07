package product

import "math/big"

type Product struct {
	Id      string
	Name    string
	Price   *big.Int
	InStock int
}
