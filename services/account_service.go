package services

import (
	"internal-transfers/models"
	"internal-transfers/repository"
)

type AccountService interface {
	CreateAccount(account models.Account) error
	GetAccount(accountID int64) (*models.Account, error)
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount(account models.Account) error {
	return s.repo.CreateAccount(account)
}

func (s *accountService) GetAccount(accountID int64) (*models.Account, error) {
	return s.repo.GetAccountByID(accountID)
}
