package handlers

import (
	"log"
	"net/http"
	"online_wallet/models"

	"github.com/gin-gonic/gin"
)

func (r Handler) ReplenishTheBalance(c *gin.Context) {
	body := models.ReplenishB{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println("Error while parse ReplenishTheBalance body")
		return
	}
	val, err := r.repo.Storage().ReplenishTheBalance(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't fall balance",
		})
	}
	c.JSON(http.StatusOK, val)
}
