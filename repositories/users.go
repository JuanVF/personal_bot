package repositories

type User struct {
	Id          int
	Name        string
	LastName    string
	GoogleMe    string
	LastUpdated string
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
