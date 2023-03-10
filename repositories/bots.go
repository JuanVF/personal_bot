package repositories

type Bot struct {
	Id            int
	UserId        int
	LastThreadId  string
	LastPaymentId int
}

// Retrieves the Bot Info for a certain user
func GetBotByUserId(userId int) (*Bot, error) {
	var bot Bot = Bot{
		UserId: userId,
	}

	statement := `SELECT
                    id,
                    last_thread_id,
                    last_payment_id
                FROM personal_bot.t_bots
                WHERE user_id = $1`

	err := db.GetConnection().QueryRow(statement, userId).Scan(
		&bot.Id,
		&bot.LastThreadId,
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
				SET last_thread_id=$1, last_payment_id=$2
				WHERE user_id = $3`

	_, err := db.GetConnection().Exec(statement, bot.LastThreadId, bot.LastPaymentId, bot.UserId)

	if err != nil {
		logger.Error("Bot Repository - Update Bot", err.Error())
	}

	return err
}
