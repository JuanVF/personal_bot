package repositories

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id          int
	Name        string
	LastName    string
	GoogleMe    string
	LastUpdated string
}

type CreateUserBody struct {
	Name     string
	LastName string
	GoogleMe string
}

type ActivityLevel struct {
	Id          int
	Name        string
	Description string
}

// Get an activity level by name
func GetActivityLevelByName(name string) (*ActivityLevel, error) {
	var activityLevel ActivityLevel

	query := "SELECT id, name, description FROM personal_bot.t_activity_levels WHERE name = $1"

	err := db.GetConnection().QueryRow(query, name).Scan(&activityLevel.Id, &activityLevel.Name, &activityLevel.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Not Found")
		} else {
			logger.Error("Activity Level Repository - Get Activity Level By Name", err.Error())

			return nil, err
		}
	}

	return &activityLevel, nil
}

// Query all the users from the DB
func GetUsers() ([]*User, error) {
	statement := "SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users"

	rows, err := db.GetConnection().Query(statement)

	if err != nil {
		logger.Error("User Repository - Get Users", err.Error())

		return []*User{}, err
	}

	defer rows.Close()

	var users []*User = make([]*User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.GoogleMe, &user.LastUpdated); err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// Returns a single user from DB
func GetUser(id int) (*User, error) {
	var user User = User{}

	statement := "SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE id = $1"

	err := db.GetConnection().QueryRow(statement, id).Scan(&user.Id, &user.Name, &user.LastName, &user.GoogleMe, &user.LastUpdated)

	if err != nil {
		logger.Error("User Repository - Get User", err.Error())

		return nil, err
	}

	return &user, nil
}

// Returns a single user from DB
func GetUserByGoogleMe(email string) (*User, error) {
	var user User = User{}

	statement := "SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE google_me = $1"

	err := db.GetConnection().QueryRow(statement, email).Scan(&user.Id, &user.Name, &user.LastName, &user.GoogleMe, &user.LastUpdated)

	if err != nil {
		logger.Error("User Repository - Get User By Google Me", err.Error())

		return nil, err
	}

	return &user, nil
}

// Creates an user in DB
func CreateUser(user *CreateUserBody) (int, error) {
	statement := `INSERT INTO personal_bot.t_users(
					name, last_name, google_me, last_updated)
				VALUES 
					($1, $2, $3, NOW()) 
				RETURNING id`

	var id int
	err := db.GetConnection().QueryRow(statement, user.Name, user.LastName, user.GoogleMe).Scan(&id)

	if err != nil {
		logger.Error("Bot Repository - CreateUser", err.Error())

		return 0, err
	}

	return id, nil
}
