package main

import (
	"internal-transfers/config"
	"internal-transfers/handlers"
	"internal-transfers/repository"
	"internal-transfers/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	config.InitDB()

	// Repositories
	accRepo := repository.NewAccountRepository(config.DB)
	txRepo := repository.NewTransactionRepository(config.DB)

	// Services
	accService := services.NewAccountService(accRepo)
	txService := services.NewTransactionService(accRepo, txRepo)

	// Handlers
	accHandler := handlers.NewAccountHandler(accService)
	txHandler := handlers.NewTransactionHandler(txService)

	// Gin Router
	r := gin.Default()

	// Routes
	r.POST("/accounts", accHandler.CreateAccount)
	r.GET("/accounts/:account_id", accHandler.GetAccount)
	r.POST("/transactions", txHandler.CreateTransaction)

	// Start server
	r.Run(":8080")
}
