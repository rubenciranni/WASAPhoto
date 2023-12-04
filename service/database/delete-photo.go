package database

func (db *appdbimpl) DeletePhoto(photoId string) error {
	_, err := db.c.Exec("DELETE FROM Photo WHERE photoId = ?", photoId)
	return err
}
