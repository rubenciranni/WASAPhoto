package database

func (db *appdbimpl) InsertFollow(followerId string, followedId string) error {
	_, err := db.c.Exec(
		"INSERT INTO Ban (followerId, followedId) VALUES (?, ?)",
		followerId, followedId)
	return err
}
