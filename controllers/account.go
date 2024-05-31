package controllers

import "accounting_transaction_api/services"

type AccountController struct {
	accountService services.AccountService
}

func CreateNewAccountController(accountService services.AccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}
