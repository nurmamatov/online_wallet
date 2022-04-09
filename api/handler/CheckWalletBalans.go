package handlers

import (
	"log"
	"net/http"
	"online_wallet/models"

	"github.com/gin-gonic/gin"
)

func (r Handler) CheckWalletBalans(c *gin.Context) {
	body := models.CheckWBM{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println("Error while parse CheckWalletBalans body")
		return
	}
	val, err := r.repo.Storage().CheckWalletBalans(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.RespW{Message: err.Error()})
	}
	c.JSON(http.StatusOK, val)
}
