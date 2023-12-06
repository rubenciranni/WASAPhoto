package database

func (db *appdbimpl) InsertBan(bannerId string, bannedId string) error {
	_, err := db.c.Exec(
		"INSERT OR IGNORE INTO Ban (bannerId, bannedId) VALUES (?, ?)",
		bannerId, bannedId)
	return err
}
