package database

func (db *appdbimpl) InsertComment(commentId string, photoId string, authorId string, text string, dateTime string) error {
	_, err := db.c.Exec(
		"INSERT INTO Comment (commentId, photoId, authorId, text, dateTime) VALUES (?, ?, ?, ?, ?)",
		commentId, photoId, authorId, text, dateTime)
	return err
}
