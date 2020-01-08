package router

import (
	"net/http"

	"github.com/akbarpambudi/shoppies/configuration"
	productDelivery "github.com/akbarpambudi/shoppies/product/delivery/http"
	"github.com/go-chi/chi"
)

type ShoppiesRouter struct {
	router       chi.Router
	routerConfig configuration.HttpConfig
}

func (r ShoppiesRouter) BindProductDelivery(productHttpDelivery *productDelivery.ProductHttpDelivery) {
	r.router.Route("/product", func(r chi.Router) {
		r.Post("/", productHttpDelivery.HandleAddProduct)
	})
}

func (r ShoppiesRouter) Serve() {
	http.ListenAndServe(r.routerConfig.Address, r.router)
}

func NewShoppiesRouter(router chi.Router, httpConfig configuration.HttpConfig) *ShoppiesRouter {
	return &ShoppiesRouter{router: router, routerConfig: httpConfig}
}
