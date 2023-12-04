package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetStream(userId string, startDate string) ([]schema.Photo, error) {
	var photoList []schema.Photo
	rows, err := db.c.Query(
		`
		SELECT Photo.*
		FROM Photo
		JOIN User ON Photo.authorId = User.userId
		JOIN Follow ON Photo.authorId = Follow.followedId
		WHERE Follow.followerId = ? AND Photo.dateTime < ?
		ORDER BY Photo.dateTime DESC
		LIMIT 20
		`,
		userId,
		startDate,
	)
	defer rows.Close()

	for rows.Next() {
		var (
			photoId          string
			authorId         string
			authorUsername   string
			caption          string
			dateTime         string
			numberOfLikes    int
			numberOfComments int
		)
		if err := rows.Scan(&photoId, &authorId, &authorUsername, &caption, &dateTime, &numberOfLikes, &numberOfComments); err != nil {
			return photoList, err
		}
		photoList = append(photoList, schema.Photo{
			PhotoId:          photoId,
			Author:           schema.User{UserId: authorId, Username: authorUsername},
			Caption:          caption,
			DateTime:         dateTime,
			NumberOfLikes:    numberOfLikes,
			NumberOfComments: numberOfComments,
		})
	}

	return photoList, err
}
