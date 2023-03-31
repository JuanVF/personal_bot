/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
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
					($1, $2, 0) 
				RETURNING id`

	_, err := db.GetConnection().Exec(statement, bot.UserId, *bot.LastGmailId)

	if err != nil {
		logger.Error("Bot Repository - CreateBot", err.Error())
	}

	return err
}
