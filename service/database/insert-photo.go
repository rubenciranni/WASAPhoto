package database

func (db *appdbimpl) InsertPhoto(photoId string, authorId string, caption string, dateTime string) error {
	_, err := db.c.Exec(
		"INSERT INTO Photo (photoId, authorId, caption, dateTime, numberOfLikes, numberOfComments) VALUES (?, ?, ?, ?, 0, 0)",
		photoId, authorId, caption, dateTime)
	return err
}
