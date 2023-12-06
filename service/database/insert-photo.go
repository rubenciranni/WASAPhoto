package database

func (db *appdbimpl) InsertPhoto(photoId string, authorId string, caption string, dateTime string) error {
	_, err := db.c.Exec(
		"INSERT INTO Photo (photoId, authorId, caption, dateTime) VALUES (?, ?, ?, ?)",
		photoId, authorId, caption, dateTime)
	return err
}
