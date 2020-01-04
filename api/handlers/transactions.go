package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/umpisa-backend/app/models"
	"github.com/jmramos02/umpisa-backend/app/services"
)

func GetUserTransactions(c *gin.Context) {
	var request services.BalanceRequest
	userContext, _ := c.Get("user")
	user, _ := userContext.(models.User)

	db, _ := c.Get("db")
	request.UserID = user.ID
	if dbObj, success := db.(*gorm.DB); success {
		response, _ := services.GetHistory(request, dbObj)
		c.JSON(200, gin.H{
			"data": response,
		})
	}

	return
}
