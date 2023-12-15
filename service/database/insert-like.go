package database

func (db *appdbimpl) InsertLike(photoID string, userID string) error {
	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO Like (photoID, userID) VALUES (?, ?)",
		photoID, userID)
	return err
}
