package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUsers(loggedInUserID string, username string, startID string) ([]schema.User, error) {
	var userList []schema.User
	rows, err := db.c.Query(
		`
		SELECT *
		FROM User
		WHERE 
			username LIKE '%' || ? || '%' COLLATE NOCASE AND
			userID > ? AND
			userID NOT IN (
				SELECT userID
				FROM BAN
				WHERE bannedId = ?
			)
		ORDER BY userID
		LIMIT 20
		`,
		username,
		startID,
		loggedInUserID,
	)
	if err != nil {
		return userList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userID   string
			username string
		)
		if err := rows.Scan(&userID, &username); err != nil {
			return userList, err
		}
		userList = append(userList, schema.User{
			UserID:   userID,
			Username: username,
		})
	}

	if err := rows.Err(); err != nil {
		return userList, err
	}

	return userList, err
}
