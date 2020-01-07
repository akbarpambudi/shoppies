package product

type ProductUseCase interface {
	FindProductById(id string) (*Product, error)
	AddProduct(productName string, stock int, price int) error
}
