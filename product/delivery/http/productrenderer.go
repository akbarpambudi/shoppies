package http

import (
	"net/http"

	"github.com/akbarpambudi/shoppies/product"
)

type GetProductRenderer struct {
	product product.Product
}

func (g GetProductRenderer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
