package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/umpisa-backend/app/models"
	"github.com/jmramos02/umpisa-backend/app/services"
	"github.com/jmramos02/umpisa-backend/app/utils"
)

func Topup(c *gin.Context) {
	var request services.TransactionRequest
	c.Bind(&request)
	userContext, _ := c.Get("user")

	user, success := userContext.(models.User)
	if !success {
		c.JSON(401, gin.H{
			"errors": "Unauthorized",
		})
	}

	request.UserID = user.ID

	db, _ := c.Get("db")
	if dbObj, success := db.(*gorm.DB); success {
		response, err := services.Topup(request, dbObj)
		if err != nil {
			if merr, ok := err.(*multierror.Error); ok {
				errors := utils.ExtractErrorMessages(merr.Errors)
				c.JSON(400, gin.H{
					"errors": errors,
				})
				return
			}
			c.JSON(400, gin.H{
				"errors": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"data": response,
		})
	}

	return
}
