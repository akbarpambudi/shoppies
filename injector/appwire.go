//+build wireinject

package injector

import (
	"github.com/akbarpambudi/shoppies/configuration"
	"github.com/akbarpambudi/shoppies/router"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

var (
	appModuleSets = wire.NewSet(configModuleSets, wire.FieldsOf(new(*configuration.AppConfig), "Http"), chi.NewRouter, wire.Bind(new(chi.Router), new(*chi.Mux)), router.NewShoppiesRouter)
)

func InitializeChiRouter() chi.Router {
	panic(wire.Build(appModuleSets))
}

func InitializeShoppiesRouter() (*router.ShoppiesRouter, error) {
	panic(wire.Build(appModuleSets))
}
