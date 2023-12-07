package database

import (
	"github.com/rubenciranni/WASAPhoto/service/model/schema"
)

func (db *appdbimpl) GetLikes(photoId string, startId string) ([]schema.User, error) {
	var userList []schema.User
	rows, err := db.c.Query(
		`
		SELECT User.* 
		FROM Like JOIN User
		ON User.userId = Like.userId
		WHERE Like.photoId = ? AND User.userId > ?
		ORDER BY User.userId
		LIMIT 20
		 `,
		photoId,
		startId,
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
