package repository

import (
	"database/sql"
	"internal-transfers/models"
)

type TransactionRepository interface {
	CreateTransaction(tx models.Transaction) error
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(tx models.Transaction) error {
	_, err := r.db.Exec(
		"INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)",
		tx.SourceAccountID, tx.DestinationAccountID, tx.Amount,
	)
	return err
}
