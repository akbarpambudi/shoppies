package main

import (
	"github.com/akbarpambudi/shoppies/injector"
	"github.com/akbarpambudi/shoppies/product/delivery/http"
	"github.com/go-chi/chi"
)

type ShoppiesRouter struct {
	router chi.Router
}

func NewShoppiesRouter(router chi.Router) *ShoppiesRouter {
	return &ShoppiesRouter{router: router}
}

func (r ShoppiesRouter) BindProductDelivery(productHttpDelivery *http.ProductHttpDelivery) {
	productDelivery := injector.InitializeProductDelivery()
	r.router.Route("/product", func(r chi.Router) {
		r.Post("/", productDelivery.HandleAddProduct)
	})
}
