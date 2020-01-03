package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmramos02/umpisa-backend/app/models"
	"github.com/jmramos02/umpisa-backend/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LoginTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SETUP TEST
func (suite *LoginTestSuite) SetupTest() {
	db, _ := gorm.Open("postgres", config.GetDatabaseUrl())
	suite.db = db

	//Remove all records
	db.Delete(&models.User{ID: 0})

	//Create a dummy user we can use for logging in
	request := RegisterRequest{
		FirstName: "JM",
		LastName:  "Ramos",
		Email:     "ramosjosemari@gmail.com",
		Password:  "testpassword",
	}

	Register(request, suite.db)
}

func (suite *LoginTestSuite) TestSuccessfulLogin() {
	request := LoginRequest{
		Email:    "ramosjosemari@gmail.com",
		Password: "testpassword",
	}

	response, err := Login(request, suite.db)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), response.AccessToken)
	assert.NotNil(suite.T(), response.User)
}

func TestLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
