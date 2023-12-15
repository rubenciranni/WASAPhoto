package database

func (db *appdbimpl) DeleteLike(photoId string, userId string) error {
	_, err := db.c.Exec("DELETE FROM Like WHERE photoId = ? AND userId = ?", photoId, userId)
	return err
}
