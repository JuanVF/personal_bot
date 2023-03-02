package repositories

import "fmt"

type Tag struct {
	Id   int
	Name string
}

// Query all the tags from the DB
func GetTags() ([]*Tag, error) {
	statement := "SELECT id, name FROM personal_bot.t_tags"

	rows, err := db.GetConnection().Query(statement)

	if err != nil {
		logger.Error("Tags Repository - Get Tags", err.Error())

		return []*Tag{}, err
	}

	defer rows.Close()

	var tags []*Tag = make([]*Tag, 0)

	for rows.Next() {
		var tag Tag

		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			return tags, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}

func insertTagPerPayment(paymentId int, tag *Tag) error {
	statement := `INSERT INTO personal_bot.t_payments_per_tags(
					id_payment, id_tag)
				VALUES ($1, $2)`

	_, err := db.GetConnection().Exec(statement, paymentId, tag.Id)

	if err != nil {
		logger.Error("Tags Repository - Insert Tag Per Payment", err.Error())

		return err
	}

	return nil
}

// Inserts the tags for a payment.
// I am using this for personal purposes, I will not bother optimizing for bulk inserts
// But if you want to, feel free to send a PR
func InsertTagsPerPayment(paymentId int, tags []*Tag) error {
	var unprocessedInserts []int = make([]int, 0)

	for _, tag := range tags {
		err := insertTagPerPayment(paymentId, tag)

		if err != nil {
			unprocessedInserts = append(unprocessedInserts, tag.Id)
		}
	}

	if len(unprocessedInserts) > 0 {
		errMsg := fmt.Sprintf("The next tags IDs couldn't be inserted for the payment Id [%d] : [%v]", paymentId, unprocessedInserts)

		logger.Error("Tags Repository - Insert Tags Per Payment", errMsg)

		return fmt.Errorf(errMsg)
	}

	return nil
}
