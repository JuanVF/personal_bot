package repositories

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var userTest *CreateUserBody = &CreateUserBody{
	Name:     "Jhon",
	LastName: "Doe",
	GoogleMe: "jhon.doe@gmail.com",
}

var fakeUserTest *CreateUserBody = &CreateUserBody{
	Name:     "Not Jhon",
	LastName: "Fake Doe",
	GoogleMe: "fake.jhon.doe@gmail.com",
}

func TestCreateUser(t *testing.T) {
	tName := "User Repository - Create User"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("INSERT INTO personal_bot.t_users( name, last_name, google_me, last_updated) VALUES ($1, $2, $3, NOW()) RETURNING id").
		WithArgs(userTest.Name, userTest.LastName, userTest.GoogleMe).
		WillReturnRows(rows)

	// Call function to be tested
	id, _ := CreateUser(userTest)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if id != 1 {
		logger.TestError(tName, "User created with id: 1", fmt.Sprintf("User created with id: %d", id), t)
	}
}

func TestGetUserByGoogleMe(t *testing.T) {
	tName := "User Repository - Get User By Google Me"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"}).AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00")

	mock.ExpectQuery("SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE google_me = $1").
		WithArgs(userTest.GoogleMe).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUserByGoogleMe(userTest.GoogleMe)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user.GoogleMe != userTest.GoogleMe {
		logger.TestError(tName, fmt.Sprintf("User with google_me: %s", userTest.GoogleMe), fmt.Sprintf("User with google_me: %s", user.GoogleMe), t)
	}

	if user.Name != userTest.Name {
		logger.TestError(tName, fmt.Sprintf("User with name: %s", userTest.Name), fmt.Sprintf("User with name: %s", user.Name), t)
	}

	if user.LastName != userTest.LastName {
		logger.TestError(tName, fmt.Sprintf("User with last name: %s", userTest.LastName), fmt.Sprintf("User with last name: %s", user.LastName), t)
	}
}

func TestGetUserByGoogleMeWhenItShouldNotExists(t *testing.T) {
	tName := "User Repository - Get User By Google Me - When it should not exists"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"})

	mock.ExpectQuery("SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE google_me = $1").
		WithArgs(fakeUserTest.GoogleMe).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUserByGoogleMe(fakeUserTest.GoogleMe)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user != nil {
		logger.TestError(tName, "User should be nil", fmt.Sprintf("User with google_me: %s", user.GoogleMe), t)
	}
}

func TestGetUser(t *testing.T) {
	tName := "User Repository - Get User"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"}).
		AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00")

	mock.ExpectQuery("SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE id = $1").
		WithArgs(1).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUser(1)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user == nil {
		logger.TestError(tName, "User should not be nil", "User is nil", t)
	}

	if user.Id != 1 {
		logger.TestError(tName, fmt.Sprintf("User id should be 1, got %d", user.Id), "", t)
	}

	if user.Name != userTest.Name {
		logger.TestError(tName, fmt.Sprintf("User with name: %s", userTest.Name), fmt.Sprintf("User with name: %s", user.Name), t)
	}

	if user.LastName != userTest.LastName {
		logger.TestError(tName, fmt.Sprintf("User with last name: %s", userTest.LastName), fmt.Sprintf("User with last name: %s", user.LastName), t)
	}

	if user.GoogleMe != userTest.GoogleMe {
		logger.TestError(tName, fmt.Sprintf("User with google_me: %s", userTest.GoogleMe), fmt.Sprintf("User with google_me: %s", user.GoogleMe), t)
	}
}

func TestGetUserWhenNotFound(t *testing.T) {
	tName := "User Repository - Get User - When not found"
	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"})

	mock.ExpectQuery("SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE id = $1").
		WithArgs(100).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUser(100)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user != nil {
		logger.TestError(tName, "User should be nil", fmt.Sprintf("User: %v", user), t)
	}
}

func TestGetAllUsers(t *testing.T) {
	tName := "User Repository - Get All Users"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"}).
		AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00").
		AddRow(2, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00")

	mock.ExpectQuery("SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users").
		WillReturnRows(rows)

	// Call function to be tested
	users, _ := GetUsers()

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if users == nil {
		logger.TestError(tName, "Users should not be nil", "Users is nil", t)
	}

	if len(users) != 2 {
		logger.TestError(tName, "Users length should be 2", fmt.Sprintf("Users length is: %d", len(users)), t)
	}
}

func TestGetAllUsersWhenNotFound(t *testing.T) {
	tName := "User Repository - Get All Users - When not found"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"})

	mock.ExpectQuery("SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users").
		WillReturnRows(rows)

	// Call function to be tested
	users, _ := GetUsers()

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if len(users) > 0 {
		logger.TestError(tName, "Users should length should be 0", fmt.Sprintf("Users: %v", users), t)
	}
}
