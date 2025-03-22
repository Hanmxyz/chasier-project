package product

type ProductRepository interface {
	GetAllProducts() ([]Product, error)
	GetProductById(Id string) (Product, error)
	CreateProduct(product Product) error
	UpdateProduct(product Product, id string) error
	DeleteProduct(id string) error
	GetProductByCategoryId(id string) ([]Product, error)
}
