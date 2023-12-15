package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUserProfile(userID string) (schema.UserProfile, error) {
	var (
		username          string
		numberOfPhotos    int
		numberOfFollowers int
		numberOfFollowing int
	)
	err := db.c.QueryRow(
		`
		SELECT 
			User.username,
			(SELECT COUNT(*) FROM Photo WHERE authorId = ?) AS numberOfPhotos,
			(SELECT COUNT(*) FROM Follow WHERE followedId = ?) AS numberOfFollowers,
			(SELECT COUNT(*) FROM Follow WHERE followerId = ?) AS numberOfFollowing
		FROM 
			User
		WHERE 
			User.userID = ?
		`,
		userID,
		userID,
		userID,
		userID).Scan(&username, &numberOfPhotos, &numberOfFollowers, &numberOfFollowing)

	userProfile := schema.UserProfile{
		UserID:            userID,
		Username:          username,
		NumberOfPhotos:    numberOfPhotos,
		NumberOfFollowers: numberOfFollowers,
		NumberOfFollowing: numberOfFollowing,
	}

	return userProfile, err
}
