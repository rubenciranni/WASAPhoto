package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetPhotosByUser(userId string, startDate string) ([]schema.Photo, error) {
	var photoList []schema.Photo
	rows, err := db.c.Query(
		`
		SELECT 
			Photo.photoId,
			User.username AS authorUsername,
			Photo.caption,
			Photo.dateTime,
			(SELECT COUNT(*) FROM Like WHERE photoId = Photo.photoId) AS numberOfLikes,
			(SELECT COUNT(*) FROM Comment WHERE photoId = Photo.photoId) AS numberOfComments
		FROM 
			Photo
		JOIN 
			User ON Photo.authorId = User.userId
		WHERE 
			Photo.authorId = ? AND Photo.dateTime < ?
		ORDER BY Photo.dateTime DESC
		LIMIT 20;
		`,
		userId,
		startDate,
	)
	defer rows.Close()

	for rows.Next() {
		var (
			photoId          string
			authorUsername   string
			caption          string
			dateTime         string
			numberOfLikes    int
			numberOfComments int
		)
		if err := rows.Scan(&photoId, &authorUsername, &caption, &dateTime, &numberOfLikes, &numberOfComments); err != nil {
			return photoList, err
		}
		photoList = append(photoList, schema.Photo{
			PhotoId:          photoId,
			Author:           schema.User{UserId: userId, Username: authorUsername},
			Caption:          caption,
			DateTime:         dateTime,
			NumberOfLikes:    numberOfLikes,
			NumberOfComments: numberOfComments,
		})
	}

	return photoList, err
}
