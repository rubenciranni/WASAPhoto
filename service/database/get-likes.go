package database

import (
	"github.com/rubenciranni/WASAPhoto/service/model/schema"
)

func (db *appdbimpl) GetLikes(photoID string, startID string) ([]schema.User, error) {
	var userList []schema.User
	rows, err := db.c.Query(
		`
		SELECT User.* 
		FROM Like JOIN User
		ON User.userID = Like.userID
		WHERE Like.photoID = ? AND User.userID > ?
		ORDER BY User.userID
		LIMIT 20
		 `,
		photoID,
		startID,
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
