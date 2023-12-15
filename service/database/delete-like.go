package database

func (db *appdbimpl) DeleteLike(photoID string, userID string) error {
	_, err := db.c.Exec("DELETE FROM Like WHERE photoID = ? AND userID = ?", photoID, userID)
	return err
}
