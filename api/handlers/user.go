package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/umpisa-backend/app/models"
)

func GetUser(c *gin.Context) {
	userContext, _ := c.Get("user")

	if user, success := userContext.(models.User); success {
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}
