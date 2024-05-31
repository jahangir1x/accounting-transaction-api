package routes

import (
	"accounting_transaction_api/controllers"
	"github.com/labstack/echo/v4"
)

type AccountRoutes struct {
	echo       *echo.Echo
	controller controllers.AccountController
}

func CreateNewAccountRoute(e *echo.Echo, controller controllers.AccountController) *AccountRoutes {
	return &AccountRoutes{
		echo:       e,
		controller: controller,
	}
}

func (ar *AccountRoutes) InitAccountRoutes() {

}
