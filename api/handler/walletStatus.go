package handlers

import (
	"log"
	"net/http"
	"online_wallet/models"

	"github.com/gin-gonic/gin"
)

func (r Handler) WalletStatus(c *gin.Context) {
	body := models.CheckWBM{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println("Error while parse WalletStatus body")
		return
	}
	val, err := r.repo.Storage().WalletStatus(&body)
	if err != nil {
		c.JSON(http.StatusOK, models.RespW{Message: "Wallet haven't"})
	}
	c.JSON(http.StatusOK, val)
}
