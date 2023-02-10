package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingResponse struct {
	Message string
}

type MainTestSuite struct {
	suite.Suite
	routerMock *gin.Engine
}

func (suite *MainTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
}

func (suite *MainTestSuite) TestPingSuccess() {
	GetPing(suite.routerMock)
	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/ping", nil)
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualResponsePing PingResponse
	json.Unmarshal([]byte(response), &actualResponsePing)
	fmt.Println("Test Struct", actualResponsePing)
	fmt.Println("Check Response", response)
	//fmt.Println("Test Message", actualResponsePing.Message)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), "Hello", actualResponsePing.Message)
	assert.Nil(suite.T(), err)
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
