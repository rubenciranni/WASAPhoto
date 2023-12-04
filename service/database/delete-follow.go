package database

func (db *appdbimpl) DeleteFollow(followerId string, followedId string) error {
	_, err := db.c.Exec("DELETE FROM Follow WHERE followerId = ? AND followedId = ?", followerId, followedId)
	return err
}
