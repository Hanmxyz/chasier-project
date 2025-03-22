package db

import (
	"chasier/domain/transaction"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) transaction.TransactionRepository {
	return TransactionRepository{DB: DB}
}

func (repository TransactionRepository) GetAllTransactions() ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	result := repository.DB.Table("transactions").Preload("Transaction_items.Product").Find(&transactions)
	return transactions, result.Error
}

func (repository TransactionRepository) CreateTransaction(items transaction.Request) error {
	var Transaction transaction.Transaction

	for _, item := range items.Items {
		Transaction.Transaction_items = append(Transaction.Transaction_items, transaction.Transaction_items{
			Product_Id: item.Product_Id,
			Amount:     item.Amount,
		})
	}

	result := repository.DB.Create(&Transaction)
	return result.Error
}
