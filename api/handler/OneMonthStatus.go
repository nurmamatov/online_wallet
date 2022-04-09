package handlers

import (
	"log"
	"net/http"
	"online_wallet/models"

	"github.com/gin-gonic/gin"
)

func (r Handler) OneMonthStatus(c *gin.Context) {
	body := models.CheckWBM{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println("Error while parse OneMonthStatus body", err)
		return
	}
	val, err := r.repo.Storage().OneMonthStatus(&body)
	if err != nil {
		c.JSON(http.StatusOK, models.RespW{Message: err.Error()})
	}
	c.JSON(http.StatusOK, val)
}
