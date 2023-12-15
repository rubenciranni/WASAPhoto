package database

func (db *appdbimpl) DeletePhoto(photoID string) error {
	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		DELETE FROM Photo WHERE photoID = ?;
		DELETE FROM Comment WHERE photoID = ?;
		DELETE FROM Like WHERE photoID = ?;
		COMMIT;
		`,
		photoID,
		photoID,
		photoID)
	return err
}
