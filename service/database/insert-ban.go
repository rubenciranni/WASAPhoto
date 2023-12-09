package database

func (db *appdbimpl) InsertBan(bannerId string, bannedId string) error {
	_, err := db.c.Exec(
		`
		BEGIN TRANSACTION;
		INSERT OR IGNORE INTO Ban (bannerId, bannedId) VALUES (?, ?);
		DELETE FROM Follow WHERE followerId = ? AND followedId = ?;
		COMMIT;
		`,
		bannerId, bannedId,
		bannedId, bannerId,
	)

	return err
}
