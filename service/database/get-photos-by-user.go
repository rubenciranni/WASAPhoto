package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetPhotosByUser(loggedInUserId string, userId string, startDate string, startId string) ([]schema.Photo, error) {
	var photoList []schema.Photo
	rows, err := db.c.Query(
		`
		SELECT 
			Photo.photoId,
			User.username AS authorUsername,
			Photo.caption,
			Photo.dateTime,
			(SELECT COUNT(*) FROM Like WHERE photoId = Photo.photoId) AS numberOfLikes,
			(SELECT COUNT(*) FROM Comment WHERE photoId = Photo.photoId) AS numberOfComments,
			(SELECT COUNT(*) FROM Like WHERE photoId = Photo.photoId AND userId = ?) AS isLiked
		FROM 
			Photo
		JOIN 
			User ON Photo.authorId = User.userId
		WHERE 
			Photo.authorId = ? AND
			Photo.dateTime < ?  OR (Photo.dateTime = ? AND Photo.photoId > ?)
		ORDER BY Photo.dateTime DESC, Photo.photoId ASC
		LIMIT 20;
		`,
		loggedInUserId,
		userId,
		startDate,
		startDate,
		startId,
	)
	if err != nil {
		return photoList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			photoId          string
			authorUsername   string
			caption          string
			dateTime         string
			numberOfLikes    int
			numberOfComments int
			isLiked          bool
		)
		if err := rows.Scan(&photoId, &authorUsername, &caption, &dateTime, &numberOfLikes, &numberOfComments, &isLiked); err != nil {
			return photoList, err
		}
		photoList = append(photoList, schema.Photo{
			PhotoId:          photoId,
			Author:           schema.User{UserId: userId, Username: authorUsername},
			Caption:          caption,
			DateTime:         dateTime,
			NumberOfLikes:    numberOfLikes,
			NumberOfComments: numberOfComments,
			IsLiked:          isLiked,
		})
	}

	if err := rows.Err(); err != nil {
		return photoList, err
	}

	return photoList, err
}
