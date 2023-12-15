package database

func (db *appdbimpl) DeletePhoto(photoId string) error {
	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		DELETE FROM Photo WHERE photoId = ?;
		DELETE FROM Comment WHERE photoId = ?;
		DELETE FROM Like WHERE photoId = ?;
		COMMIT;
		`,
		photoId,
		photoId,
		photoId)
	return err
}
