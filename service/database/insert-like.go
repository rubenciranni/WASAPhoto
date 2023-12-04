package database

func (db *appdbimpl) InsertLike(photoId string, userId string) error {
	_, err := db.c.Exec(
		"INSERT INTO Like (photoId, userId) VALUES (?, ?)",
		photoId, userId)
	return err
}
