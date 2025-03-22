package db

import (
	"chasier/domain/categories"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) categories.CategoryRepository {
	return CategoryRepository{DB: DB}
}

func (repository CategoryRepository) GetAllCategories() ([]categories.Category, error) {
	var categories []categories.Category
	result := repository.DB.Table("categories").Find(&categories)
	return categories, result.Error
}

func (repository CategoryRepository) CreateCategory(category categories.Category) error {
	result := repository.DB.Table("categories").Create(&category)
	return result.Error
}

func (repository CategoryRepository) UpdateCategory(category categories.Category, id string) error {
	result := repository.DB.Table("categories").Where("id = ?", id).Updates(&category)
	return result.Error
}

func (repository CategoryRepository) DeleteCategory(id string) error {
	result := repository.DB.Table("categories").Where("id = ?", id).Delete(categories.Category{})
	return result.Error
}
