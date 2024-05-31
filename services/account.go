package services

import "accounting_transaction_api/repositories"

type AccountService struct {
	accountRepo repositories.AccountRepo
}

func CreateNewAccountService(accountRepo repositories.AccountRepo) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}
