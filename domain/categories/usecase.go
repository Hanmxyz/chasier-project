package categories

type CategoryUsecase struct {
	Repository CategoryRepository
}

func (Usecase CategoryUsecase) GetAllCategories() ([]Category, error) {
	Categories, err := Usecase.Repository.GetAllCategories()
	return Categories, err
}

func (Usecase CategoryUsecase) CreateCategory(category Category) error {
	err := Usecase.Repository.CreateCategory(category)
	return err
}

func (Usecase CategoryUsecase) UpdateCategory(category Category, id string) error {
	err := Usecase.Repository.UpdateCategory(category, id)
	return err
}

func (usecase CategoryUsecase) DeleteCategory(id string) error {
	err := usecase.Repository.DeleteCategory(id)
	return err
}
