package containers

import (
	"accounting_transaction_api/config"
	"accounting_transaction_api/connection"
	"accounting_transaction_api/controllers"
	"accounting_transaction_api/repositories"
	"accounting_transaction_api/routes"
	"accounting_transaction_api/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func Serve(e *echo.Echo) {
	config.SetConfig()
	db := connection.GetDB()

	// repository initialization
	accountRepo := repositories.CreateNewAccountRepository(db)

	// service initialization
	accountService := services.CreateNewAccountService(*accountRepo)

	// controller initialization
	accountCtr := controllers.CreateNewAccountController(*accountService)

	// route initialization
	accountRoutes := routes.CreateNewAccountRoute(e, *accountCtr)

	// route binding
	accountRoutes.InitAccountRoutes()

	// start server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
