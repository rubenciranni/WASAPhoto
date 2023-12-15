package database

func (db *appdbimpl) InsertComment(commentID string, photoID string, authorId string, text string, dateTime string) error {
	_, err := db.c.Exec(
		"INSERT INTO Comment (commentID, photoID, authorId, text, dateTime) VALUES (?, ?, ?, ?, ?)",
		commentID, photoID, authorId, text, dateTime)
	return err
}
