package database

func (db *appdbimpl) InsertUser(userId string, username string) error {
	_, err := db.c.Exec(
		`
		INSERT INTO User (userId, username)
		VALUES (?, ?)
		`,
		userId,
		username)
	return err
}
