package database

func (db *appdbimpl) ExistsBan(bannerId string, bannedId string) (bool, error) {
	var exists bool
	err := db.c.QueryRow("SELECT COUNT(*) FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId).Scan(&exists)
	return exists, err
}
