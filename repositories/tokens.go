package repositories

type Token struct {
	Id           int
	User         *User
	RefreshToken string
	LastToken    string
	ExpiresIn    string
	LastUpdated  string
}

// Retrieves the Token Info for a certain user
func GetTokenByUserId(userId int) (*Token, error) {
	var user User = User{}
	var token Token = Token{
		User: &user,
	}

	statement := `SELECT 
                    tok.id token_id, 
                    tok.refresh_token, 
                    tok.last_token, 
                    tok.expires_in, 
                    tok.last_updated, 
                    tok.user_id user_id,
                    tus.name,
                    tus.last_name,
                    tus.google_me,
                    tus.last_updated
                FROM personal_bot.t_tokens tok
                INNER JOIN personal_bot.t_users tus
                    ON tok.user_id = tus.id
                WHERE user_id = $1`

	err := db.GetConnection().QueryRow(statement, userId).Scan(
		&token.Id,
		&token.RefreshToken,
		&token.LastToken,
		&token.ExpiresIn,
		&token.LastUpdated,
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.GoogleMe,
		&user.LastUpdated,
	)

	if err != nil {
		logger.Error("Token Repository - Get Token By User Id", err.Error())

		return nil, err
	}

	return &token, nil
}

// Inserts a token into the db and sets the Id to the token variable
func InsertToken(token *Token) error {
	statement := `INSERT INTO personal_bot.t_tokens(
					refresh_token, last_token, expires_in, last_updated, user_id)
				VALUES 
					($1, $2, $3, NOW(), $4) 
				RETURNING id`

	err := db.GetConnection().QueryRow(statement, token.RefreshToken, token.LastToken, token.ExpiresIn, token.User.Id).Scan(&token.Id)

	if err != nil {
		logger.Error("Token Repository - Insert Token", err.Error())
	}

	return err
}

// Updates a Token info
func UpdateToken(token *Token) error {
	statement := `UPDATE personal_bot.t_tokens
				SET refresh_token=$1, last_token=$2, expires_in=$3, last_updated=NOW()
				WHERE user_id = $4`

	_, err := db.GetConnection().Exec(statement, token.RefreshToken, token.LastToken, token.ExpiresIn, token.User.Id)

	if err != nil {
		logger.Error("Token Repository - Update Token", err.Error())
	}

	return err
}
