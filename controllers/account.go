package controllers

import (
	"accounting_transaction_api/services"
	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountService services.AccountService
}

func CreateNewAccountController(accountService services.AccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

func (controller *AccountController) GetTransactions(context echo.Context) error {
	accountInfo, err := controller.accountService.GetTransactions()
	if err != nil {
		return context.JSON(500, err.Error())
	}
	return context.JSON(200, accountInfo)
}
