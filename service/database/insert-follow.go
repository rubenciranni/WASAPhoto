package database

func (db *appdbimpl) InsertFollow(followerId string, followedId string) error {
	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO Follow (followerId, followedId) VALUES (?, ?)",
		followerId, followedId)
	return err
}
