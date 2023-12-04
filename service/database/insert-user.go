package database

func (db *appdbimpl) InsertUser(userId string, username string) error {
	_, err := db.c.Exec(
		`
		INSERT INTO User (userId, username, numberOfPhotos, numberOfFollowers, numberOfFollowing)
		VALUES (?, ?, 0, 0, 0)
		`,
		userId,
		username)
	return err
}
