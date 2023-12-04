package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetFollowing(userId string, startId string) ([]schema.User, error) {
	var followingList []schema.User
	rows, err := db.c.Query(
		`
		SELECT Follow.followedId, User.username
		FROM Follow JOIN User
		ON Follow.followedId = User.userId
		WHERE Follow.followerId = ? AND Follow.followedId > ?
		ORDER BY Follow.followedId
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
			return followingList, err
		}
		followingList = append(followingList, schema.User{
			UserId:   userId,
			Username: username,
		})
	}

	return followingList, err
}
