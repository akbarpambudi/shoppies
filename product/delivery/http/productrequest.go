package http

type AddProductRequest struct {
	ProductName string `json:"productName"`
	Stock       int
	Price       int
}
