package repositories

import (
	"fmt"
	"testing"

	"github.com/JuanVF/personal_bot/common"
)

// TestMain is the main function for the tests in this package
func TestMain(m *testing.M) {
	db := common.GetDB()
	conn := db.GetConnection()

	if conn == nil {
		err := fmt.Sprintf("%s %s", "Test DB Initialized", "Test DB is nil")

		logger.Error("Repositories - Test Main", err)
		panic(err)
	}

	mock := db.GetMock()

	if mock == nil {
		err := fmt.Sprintf("%s %s", "Test Mock Initialized", "Test Mock is nil")

		logger.Error("Repositories - Test Main", err)
		panic(err)
	}

	m.Run()
}
