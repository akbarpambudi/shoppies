package http

import (
	"encoding/json"
	"net/http"

	"github.com/akbarpambudi/shoppies/product"
)

type ProductHttpDelivery struct {
	useCase product.ProductUseCase
}

func (p *ProductHttpDelivery) HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	requestBody := new(AddProductRequest)
	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		panic(err)
	}
	if err := p.useCase.AddProduct(requestBody.ProductName, requestBody.Stock, requestBody.Price); err != nil {
		panic(err)
	}
	w.WriteHeader(201)
}

func NewProductHttpDelivery(useCase product.ProductUseCase) *ProductHttpDelivery {
	return &ProductHttpDelivery{useCase: useCase}
}
