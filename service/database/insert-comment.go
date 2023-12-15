package database

func (db *appdbimpl) InsertComment(commentId string, photoID string, authorId string, text string, dateTime string) error {
	_, err := db.c.Exec(
		"INSERT INTO Comment (commentId, photoID, authorId, text, dateTime) VALUES (?, ?, ?, ?, ?)",
		commentId, photoID, authorId, text, dateTime)
	return err
}
