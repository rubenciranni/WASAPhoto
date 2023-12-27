package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetFollowers(userId string, startId string) ([]schema.User, error) {
	var followersList []schema.User
	rows, err := db.c.Query(
		`
		SELECT User.*
		FROM Follow JOIN User
		ON Follow.followerId = User.userId
		WHERE Follow.followedId = ? AND Follow.followerId > ?
		ORDER BY Follow.followerId
		LIMIT 20
		`,
		userId,
		startId,
	)
	if err != nil {
		return followersList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userId   string
			username string
		)
		if err := rows.Scan(&userId, &username); err != nil {
			return followersList, err
		}
		followersList = append(followersList, schema.User{
			UserId:   userId,
			Username: username,
		})
	}

	if err := rows.Err(); err != nil {
		return followersList, err
	}

	return followersList, err
}
