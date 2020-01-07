//+build wireinject

package injector

import (
	"github.com/akbarpambudi/shoppies/product"
	productDeliveryHttp "github.com/akbarpambudi/shoppies/product/delivery/http"
	productRepository "github.com/akbarpambudi/shoppies/product/repository"
	productUseCase "github.com/akbarpambudi/shoppies/product/usecase"
	"github.com/google/wire"
)

var (
	productServiceSets = wire.NewSet(
		productRepository.NewInMemoryProductRepository,
		wire.Bind(new(product.ProductRepository), new(*productRepository.InMemoryProductRepository)),
		productUseCase.NewProductUseCaseImpl,
		wire.Bind(new(product.ProductUseCase), new(*productUseCase.ProductUseCaseImpl)),
	)
)

func InitializeProductDelivery() *productDeliveryHttp.ProductHttpDelivery {
	panic(wire.Build(productDeliveryHttp.NewProductHttpDelivery, productServiceSets))
}
