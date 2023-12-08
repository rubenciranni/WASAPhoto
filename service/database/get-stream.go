package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetStream(userId string, startDate string, startId string) ([]schema.Photo, error) {
	var photoList []schema.Photo
	rows, err := db.c.Query(
		`
		SELECT 
			Photo.photoId,
			Photo.authorId,
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
			Photo.authorId IN (
				SELECT followedId FROM Follow WHERE followerId = ?
			) AND
			Photo.dateTime < ?  OR (Photo.dateTime = ? AND Photo.photoId > ?)
		ORDER BY Photo.dateTime DESC, Photo.photoId ASC
		LIMIT 20
		`,
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

	if err := rows.Err(); err != nil {
		return photoList, err
	}

	return photoList, err
}
