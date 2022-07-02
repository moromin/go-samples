package sample_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MySuite struct {
	suite.Suite
	count int
}

// Suite Interfaces
func (suite *MySuite) SetupSuite() {
	suite.count++
	suite.T().Log(suite.count, ": SetupSuite")
}
func (suite *MySuite) TearDownSuite() {
	suite.count++
	suite.T().Log(suite.count, ": TearDownSuite")
}
func (suite *MySuite) SetupTest() {
	suite.count++
	suite.T().Log(suite.count, ": SetupTest")
}
func (suite *MySuite) TearDownTest() {
	suite.count++
	suite.T().Log(suite.count, ": TeardownTest")
}
func (suite *MySuite) BeforeTest(suiteName, testName string) {
	suite.count++
	suite.T().Log(suite.count, ": BeforeTest")
}
func (suite *MySuite) AfterTest(suiteName, testName string) {
	suite.count++
	suite.T().Log(suite.count, ": AfterTest")
}

func (suite *MySuite) TestSuccess() {
	suite.Assert().Equal(1, 1)
}

func (suite *MySuite) TestFail() {
	suite.Assert().NotEqual(1, 2)
}

func TestMySuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}
