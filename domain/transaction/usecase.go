package transaction

type TransactionUsecase struct {
	Repository TransactionRepository
}

func (usecase TransactionUsecase) GetAllTransactions() ([]Transaction, error) {
	transactions, err := usecase.Repository.GetAllTransactions()
	return transactions, err
}

func (usecase TransactionUsecase) CreateTransaction(items Request) error {

	err := usecase.Repository.CreateTransaction(items)
	return err

}
