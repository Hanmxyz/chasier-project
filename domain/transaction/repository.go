package transaction

type TransactionRepository interface {
	GetAllTransactions() ([]Transaction, error)
	CreateTransaction(items Request) error
}
