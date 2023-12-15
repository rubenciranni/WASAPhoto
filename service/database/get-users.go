package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUsers(loggedInUserId string, username string, startId string) ([]schema.User, error) {
	var userList []schema.User
	rows, err := db.c.Query(
		`
		SELECT *
		FROM User
		WHERE 
			username LIKE '%' || ? || '%' COLLATE NOCASE AND
			userId > ? AND
			userId NOT IN (
				SELECT userId
				FROM BAN
				WHERE bannedId = ?
			)
		ORDER BY userId
		LIMIT 20
		`,
		username,
		startId,
		loggedInUserId,
	)
	if err != nil {
		return userList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userId   string
			username string
		)
		if err := rows.Scan(&userId, &username); err != nil {
			return userList, err
		}
		userList = append(userList, schema.User{
			UserId:   userId,
			Username: username,
		})
	}

	if err := rows.Err(); err != nil {
		return userList, err
	}

	return userList, err
}
