package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUserProfile(userId string) (schema.UserProfile, error) {
	var (
		username          string
		numberOfPhotos    int
		numberOfFollowers int
		numberOfFollowing int
	)
	err := db.c.QueryRow(
		`SELECT * FROM User WHERE userId = ?`,
		userId).Scan(&userId, &username, &numberOfPhotos, &numberOfFollowers, &numberOfFollowing)

	userProfile := schema.UserProfile{
		UserId:            userId,
		Username:          username,
		NumberOfPhotos:    numberOfPhotos,
		NumberOfFollowers: numberOfFollowers,
		NumberOfFollowing: numberOfFollowing,
	}

	return userProfile, err
}
