package database

func (db *appdbimpl) InsertPhoto(photoID string, authorId string, caption string, dateTime string) error {
	_, err := db.c.Exec(
		"INSERT INTO Photo (photoID, authorId, caption, dateTime) VALUES (?, ?, ?, ?)",
		photoID, authorId, caption, dateTime)
	return err
}
