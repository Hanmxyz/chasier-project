package db

import (
	"chasier/domain/product"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) product.ProductRepository {
	return ProductRepository{DB: DB}
}

func (repository ProductRepository) GetAllProducts() ([]product.Product, error) {
	var products []product.Product
	result := repository.DB.Table("products").Joins("Category").Find(&products)
	return products, result.Error
}

func (repository ProductRepository) GetProductById(Id string) (product.Product, error) {
	var product product.Product
	result := repository.DB.Table("products").Joins("Category").Where("products.id = ?", Id).Find(&product)
	return product, result.Error
}

func (repository ProductRepository) CreateProduct(product product.Product) error {
	result := repository.DB.Table("products").Create(&product)
	return result.Error
}

func (repository ProductRepository) UpdateProduct(product product.Product, id string) error {
	result := repository.DB.Table("products").Where("id = ?", id).Updates(&product)
	return result.Error
}

func (repository ProductRepository) DeleteProduct(id string) error {
	result := repository.DB.Table("products").Where("id = ?", id).Delete(product.Product{})
	return result.Error

}

func (repository ProductRepository) GetProductByCategoryId(id string) ([]product.Product, error) {
	var products []product.Product
	ressult := repository.DB.Table("products").Where("category_id = ?", id).Scan(&products)
	return products, ressult.Error

}
