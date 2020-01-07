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
	shoppiesRouter := NewShoppiesRouter(r)
	shoppiesRouter.BindProductDelivery(injector.InitializeProductDelivery())
	http.ListenAndServe(config.Http.Address, r)
}
