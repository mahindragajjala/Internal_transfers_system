package services

import (
	"errors"
	"internal-transfers/models"
	"internal-transfers/repository"
)

type TransactionService interface {
	ProcessTransaction(tx models.Transaction) error
}

type transactionService struct {
	accRepo repository.AccountRepository
	txRepo  repository.TransactionRepository
}

func NewTransactionService(accRepo repository.AccountRepository, txRepo repository.TransactionRepository) TransactionService {
	return &transactionService{accRepo: accRepo, txRepo: txRepo}
}

func (s *transactionService) ProcessTransaction(tx models.Transaction) error {
	sourceAcc, err := s.accRepo.GetAccountByID(tx.SourceAccountID)
	if err != nil {
		return errors.New("source account not found")
	}

	destAcc, err := s.accRepo.GetAccountByID(tx.DestinationAccountID)
	if err != nil {
		return errors.New("destination account not found")
	}

	if sourceAcc.Balance < tx.Amount {
		return errors.New("insufficient balance")
	}

	// Update balances
	s.accRepo.UpdateAccountBalance(sourceAcc.AccountID, sourceAcc.Balance-tx.Amount)
	s.accRepo.UpdateAccountBalance(destAcc.AccountID, destAcc.Balance+tx.Amount)

	// Record transaction
	return s.txRepo.CreateTransaction(tx)
}
