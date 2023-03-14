package repositories

type Bot struct {
	Id            int
	UserId        int
	LastGmailId   *string
	LastPaymentId int
}

type CreateBotBody struct {
	UserId      int
	LastGmailId *string
}

// Retrieves the Bot Info for a certain user
func GetBotByUserId(userId int) (*Bot, error) {
	var bot Bot = Bot{
		UserId: userId,
	}

	statement := `SELECT
                    id,
                    last_gmail_id,
                    last_payment_id
                FROM personal_bot.t_bots
                WHERE user_id = $1`

	err := db.GetConnection().QueryRow(statement, userId).Scan(
		&bot.Id,
		&bot.LastGmailId,
		&bot.LastPaymentId,
	)

	if err != nil {
		logger.Error("Bot Repository - Get Bot By User Id", err.Error())

		return nil, err
	}

	return &bot, nil
}

// Updates a bot info
func UpdateBot(bot *Bot) error {
	statement := `UPDATE personal_bot.t_bots
				SET last_gmail_id=$1, last_payment_id=$2
				WHERE user_id = $3`

	_, err := db.GetConnection().Exec(statement, bot.LastGmailId, bot.LastPaymentId, bot.UserId)

	if err != nil {
		logger.Error("Bot Repository - Update Bot", err.Error())
	}

	return err
}

// Updates a bot info
func CreateBot(bot *CreateBotBody) error {
	statement := `INSERT INTO personal_bot.t_bots(
					user_id, last_gmail_id, last_payment_id)
				VALUES 
					($1, $2, $3) 
				RETURNING id`

	_, err := db.GetConnection().Exec(statement, bot.UserId, bot.LastGmailId, 0)

	if err != nil {
		logger.Error("Bot Repository - CreateBot", err.Error())
	}

	return err
}