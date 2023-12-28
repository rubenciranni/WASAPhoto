package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUsers(loggedInUserId string, username string, startId string) ([]schema.User, error) {
	var userList []schema.User
	rows, err := db.c.Query(
		`
		SELECT *
		FROM User
		WHERE 
			User.username LIKE '%' || ? || '%' COLLATE NOCASE AND
			User.userId > ? AND
			NOT EXISTS (
				SELECT 1
				FROM BAN
				WHERE bannerId = User.userId AND bannedId = ?
			)
		ORDER BY User.userId
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
