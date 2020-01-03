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

type RegisterTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SETUP TEST
func (suite *RegisterTestSuite) SetupTest() {
	db, _ := gorm.Open("postgres", config.GetDatabaseUrl())
	suite.db = db

	//Remove all records
	db.Delete(&models.User{ID: 0})
}

func (suite *RegisterTestSuite) TestSuccessfulRegistration() {
	request := RegisterRequest{
		FirstName: "JM",
		LastName:  "Ramos",
		Email:     "ramosjosemari@gmail.com",
		Password:  "testpassword",
	}

	response, err := Register(request, suite.db)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), response)
	assert.NotEqual(suite.T(), 0, response.User.ID)
}

func TestRegistertSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
