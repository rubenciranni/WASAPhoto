package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetFollowers(userID string, startID string) ([]schema.User, error) {
	var followersList []schema.User
	rows, err := db.c.Query(
		`
		SELECT User.*
		FROM Follow JOIN User
		ON Follow.followerId = User.userID
		WHERE Follow.followedId = ? AND Follow.followerId > ?
		ORDER BY startID
		LIMIT 20
		`,
		userID,
		startID,
	)
	if err != nil {
		return followersList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userID   string
			username string
		)
		if err := rows.Scan(&userID, &username); err != nil {
			return followersList, err
		}
		followersList = append(followersList, schema.User{
			UserID:   userID,
			Username: username,
		})
	}

	if err := rows.Err(); err != nil {
		return followersList, err
	}

	return followersList, err
}
