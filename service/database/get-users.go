package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUsers(username string, startId string) ([]schema.User, error) {
	var userList []schema.User
	rows, err := db.c.Query(
		`
		SELECT userId, username
		FROM User
		WHERE username LIKE '%?%' AND userId > ?
		ORDER BY userId
		LIMIT 20
		`,
		username,
		startId,
	)
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

	return userList, err
}
