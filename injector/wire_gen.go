// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/akbarpambudi/shoppies/configuration"
	"github.com/akbarpambudi/shoppies/product"
	"github.com/akbarpambudi/shoppies/product/delivery/http"
	"github.com/akbarpambudi/shoppies/product/repository"
	"github.com/akbarpambudi/shoppies/product/usecase"
	"github.com/akbarpambudi/shoppies/router"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

// Injectors from appwire.go:

func InitializeChiRouter() chi.Router {
	mux := chi.NewRouter()
	return mux
}

func InitializeShoppiesRouter() (*router.ShoppiesRouter, error) {
	mux := chi.NewRouter()
	environmentConfigBinderProperties := _wireEnvironmentConfigBinderPropertiesValue
	appConfig, err := provideAppConfig(environmentConfigBinderProperties)
	if err != nil {
		return nil, err
	}
	httpConfig := appConfig.Http
	shoppiesRouter := router.NewShoppiesRouter(mux, httpConfig)
	return shoppiesRouter, nil
}

var (
	_wireEnvironmentConfigBinderPropertiesValue = configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}
)

// Injectors from configwire.go:

func InitializeAppConfig() (*configuration.AppConfig, error) {
	environmentConfigBinderProperties := _wireEnvironmentConfigBinderPropertiesValue
	appConfig, err := provideAppConfig(environmentConfigBinderProperties)
	if err != nil {
		return nil, err
	}
	return appConfig, nil
}

// Injectors from productwire.go:

func InitializeProductDelivery() *http.ProductHttpDelivery {
	inMemoryProductRepository := repository.NewInMemoryProductRepository()
	productUseCaseImpl := usecase.NewProductUseCaseImpl(inMemoryProductRepository)
	productHttpDelivery := http.NewProductHttpDelivery(productUseCaseImpl)
	return productHttpDelivery
}

// appwire.go:

var (
	appModuleSets = wire.NewSet(configModuleSets, wire.FieldsOf(new(*configuration.AppConfig), "Http"), chi.NewRouter, wire.Bind(new(chi.Router), new(*chi.Mux)), router.NewShoppiesRouter)
)

// configwire.go:

var (
	configModuleSets = wire.NewSet(wire.Value(configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}), provideAppConfig)
)

func provideAppConfig(properties configuration.EnvironmentConfigBinderProperties) (*configuration.AppConfig, error) {
	environmentConfigBinder := configuration.NewEnvironmentConfigBinder(properties)
	environmentConfigBinder.Bind()
	return environmentConfigBinder.GetAppConfig()

}

// productwire.go:

var (
	productModuleSets = wire.NewSet(repository.NewInMemoryProductRepository, wire.Bind(new(product.ProductRepository), new(*repository.InMemoryProductRepository)), usecase.NewProductUseCaseImpl, wire.Bind(new(product.ProductUseCase), new(*usecase.ProductUseCaseImpl)), http.NewProductHttpDelivery)
)
