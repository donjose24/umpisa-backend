package services

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/umpisa-backend/app/models"
	"github.com/jmramos02/umpisa-backend/app/utils"
	"gopkg.in/go-playground/validator.v9"
)

type BalanceRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}

type TransactionResponse struct {
	Status string `json:"status"`
}

type TransactionRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	Amount int  `json:"amount" validate:"required,min=200"`
}

type BalanceResponse struct {
	Amount uint `json:"amount"`
}

func Topup(request TransactionRequest, db *gorm.DB) (TransactionResponse, error) {
	v := validator.New()
	err := v.Struct(request)
	var result error

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			result = multierror.Append(result, errors.New(utils.FormatErrors(e.ActualTag(), e.Field(), e.Param())))
		}

		return TransactionResponse{}, result
	}

	transaction := models.Transaction{
		Amount: request.Amount,
		UserID: request.UserID,
	}

	db.Create(&transaction)

	return TransactionResponse{Status: "ok"}, nil
}

func Balance(request BalanceRequest, db *gorm.DB) (BalanceResponse, error) {
	var balance BalanceResponse
	db.Table("transactions").Select("sum(amount) as amount").Where("user_id = ?", request.UserID).Scan(&balance)

	return balance, nil
}
