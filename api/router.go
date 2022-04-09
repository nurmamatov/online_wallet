package api

import (
	"online_wallet/api/handlers"
	"online_wallet/config"
	"online_wallet/storage"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Conf config.Config
	Repo storage.Storage
}

func NewRouter(option Option) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	handler := handlers.NewHandler(option.Conf, option.Repo)

	router.POST("/check/wallet", handler.WalletStatus)
	router.POST("/balance/wallet", handler.CheckWalletBalans)
	router.POST("/fell/wallet", handler.ReplenishTheBalance)
	router.POST("/month/wallet", handler.OneMonthStatus)

	return router
}
