package categories

type CategoryRepository interface {
	GetAllCategories() ([]Category, error)
	CreateCategory(category Category) error
	UpdateCategory(category Category, id string) error
	DeleteCategory(id string) error
}
