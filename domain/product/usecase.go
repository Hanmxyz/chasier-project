package product

type ProductUsecase struct {
	Repository ProductRepository
}

func (usecase ProductUsecase) GetAllProducts() ([]Product, error) {
	products, err := usecase.Repository.GetAllProducts()
	return products, err
}

func (usecase ProductUsecase) GetProductById(Id string) (Product, error) {
	product, err := usecase.Repository.GetProductById(Id)
	return product, err
}

func (usecase ProductUsecase) CreateProduct(product Product) error {
	err := usecase.Repository.CreateProduct(product)
	return err
}

func (usecase ProductUsecase) UpdateProduct(product Product, id string) error {
	err := usecase.Repository.UpdateProduct(product, id)
	return err
}

func (usecase ProductUsecase) DeleteProduct(id string) error {
	err := usecase.Repository.DeleteProduct(id)
	return err
}

func (usecase ProductUsecase) GetProductByCategoryId(id string) ([]Product, error) {
	products, err := usecase.Repository.GetProductByCategoryId(id)
	return products, err

}
