package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetStream(userID string, startDate string, startId string) ([]schema.Photo, error) {
	var photoList []schema.Photo
	rows, err := db.c.Query(
		`
		SELECT 
			Photo.photoID,
			Photo.authorId,
			User.username AS authorUsername,
			Photo.caption,
			Photo.dateTime,
			(SELECT COUNT(*) FROM Like WHERE photoID = Photo.photoID) AS numberOfLikes,
			(SELECT COUNT(*) FROM Comment WHERE photoID = Photo.photoID) AS numberOfComments
		FROM 
			Photo
		JOIN 
			User ON Photo.authorId = User.userID
		WHERE 
			Photo.authorId IN (
				SELECT followedId FROM Follow WHERE followerId = ?
			) AND
			Photo.dateTime < ?  OR (Photo.dateTime = ? AND Photo.photoID > ?)
		ORDER BY Photo.dateTime DESC, Photo.photoID ASC
		LIMIT 20
		`,
		userID,
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
			photoID          string
			authorId         string
			authorUsername   string
			caption          string
			dateTime         string
			numberOfLikes    int
			numberOfComments int
		)
		if err := rows.Scan(&photoID, &authorId, &authorUsername, &caption, &dateTime, &numberOfLikes, &numberOfComments); err != nil {
			return photoList, err
		}
		photoList = append(photoList, schema.Photo{
			PhotoID:          photoID,
			Author:           schema.User{UserID: authorId, Username: authorUsername},
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
