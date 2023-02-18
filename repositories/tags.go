package repositories

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
