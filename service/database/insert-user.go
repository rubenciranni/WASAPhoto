package database

func (db *appdbimpl) InsertUser(userID string, username string) error {
	_, err := db.c.Exec(
		`
		INSERT INTO User (userID, username)
		VALUES (?, ?)
		`,
		userID,
		username)
	return err
}
