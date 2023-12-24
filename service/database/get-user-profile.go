package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUserProfile(loggedInUserId string, userId string) (schema.UserProfile, error) {
	var (
		username          string
		numberOfPhotos    int
		numberOfFollowers int
		numberOfFollowing int
		isFollowed        bool
	)
	err := db.c.QueryRow(
		`
		SELECT 
			User.username,
			(SELECT COUNT(*) FROM Photo WHERE authorId = ?) AS numberOfPhotos,
			(SELECT COUNT(*) FROM Follow WHERE followedId = ?) AS numberOfFollowers,
			(SELECT COUNT(*) FROM Follow WHERE followerId = ?) AS numberOfFollowing,
			(SELECT COUNT(*) FROM Follow WHERE followerId = ? AND followedId = ?) AS isFollowed
		FROM 
			User
		WHERE 
			User.userId = ?
		`,
		userId,
		userId,
		userId,
		loggedInUserId, userId,
		userId).Scan(&username, &numberOfPhotos, &numberOfFollowers, &numberOfFollowing, &isFollowed)

	userProfile := schema.UserProfile{
		UserId:            userId,
		Username:          username,
		NumberOfPhotos:    numberOfPhotos,
		NumberOfFollowers: numberOfFollowers,
		NumberOfFollowing: numberOfFollowing,
		IsFollowed:        isFollowed,
	}

	return userProfile, err
}
