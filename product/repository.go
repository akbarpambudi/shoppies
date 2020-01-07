package product

type ProductRepository interface {
	FindById(id string) (*Product, error)
	Add(product Product) error
}
