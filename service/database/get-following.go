package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetFollowing(userID string, startID string) ([]schema.User, error) {
	var followingList []schema.User
	rows, err := db.c.Query(
		`
		SELECT User.*
		FROM Follow JOIN User
		ON Follow.followedId = User.userID
		WHERE Follow.followerId = ? AND Follow.followedId > ?
		ORDER BY Follow.followedId
		LIMIT 20
		`,
		userID,
		startID,
	)
	if err != nil {
		return followingList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userID   string
			username string
		)
		if err := rows.Scan(&userID, &username); err != nil {
			return followingList, err
		}
		followingList = append(followingList, schema.User{
			UserID:   userID,
			Username: username,
		})
	}

	if err := rows.Err(); err != nil {
		return followingList, err
	}

	return followingList, err
}
