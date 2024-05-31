package main

import (
	"accounting_transaction_api/containers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Serve(e)
}
