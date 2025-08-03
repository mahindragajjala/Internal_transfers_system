package repository

import (
	"database/sql"
	"errors"
	"internal-transfers/models"
)

type AccountRepository interface {
	CreateAccount(account models.Account) error
	GetAccountByID(accountID int64) (*models.Account, error)
	UpdateAccountBalance(accountID int64, newBalance float64) error
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(account models.Account) error {
	_, err := r.db.Exec("INSERT INTO accounts (account_id, balance) VALUES ($1, $2)", account.AccountID, account.Balance)
	return err
}

func (r *accountRepository) GetAccountByID(accountID int64) (*models.Account, error) {
	var acc models.Account
	row := r.db.QueryRow("SELECT account_id, balance FROM accounts WHERE account_id=$1", accountID)
	if err := row.Scan(&acc.AccountID, &acc.Balance); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}
		return nil, err
	}
	return &acc, nil
}

func (r *accountRepository) UpdateAccountBalance(accountID int64, newBalance float64) error {
	_, err := r.db.Exec("UPDATE accounts SET balance=$1 WHERE account_id=$2", newBalance, accountID)
	return err
}
