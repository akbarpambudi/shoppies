package main

import (
	"net/http"

	"github.com/akbarpambudi/shoppies/injector"
	"github.com/go-chi/chi"
)

func main() {
	configBinder := injector.InitializeAppConfigBinder()
	configBinder.Bind()
	config, err := configBinder.GetAppConfig()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	productDelivery := injector.InitializeProductDelivery()
	r.Route("/product", func(r chi.Router) {
		r.Post("/", productDelivery.HandleAddProduct)
	})

	http.ListenAndServe(config.Http.Address, r)
}
