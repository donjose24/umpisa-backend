package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/umpisa-backend/app/models"
	"github.com/jmramos02/umpisa-backend/app/services"
)

func GetUserBalance(c *gin.Context) {
	var request services.TransactionRequest
	userContext, _ := c.Get("user")
	user, _ := userContext.(models.User)

	db, _ := c.Get("db")
	request.UserID = user.ID
	if dbObj, success := db.(*gorm.DB); success {
		response, _ := services.Topup(request, dbObj)
		c.JSON(200, gin.H{
			"data": response,
		})
	}

	return
}
