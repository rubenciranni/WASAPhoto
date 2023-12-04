package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetFollowers(userId string, startId string) ([]schema.User, error) {
	var followersList []schema.User
	rows, err := db.c.Query(
		`
		SELECT Follow.followerId, User.username
		FROM Follow JOIN User
		ON Follow.followerId = User.userId
		WHERE Follow.followedId = ? AND Follow.followerId > ?
		ORDER BY startId
		LIMIT 20
		`,
		userId,
		startId,
	)
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

	return followersList, err
}
