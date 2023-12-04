package database

func (db *appdbimpl) DeleteBan(bannerId string, bannedId string) error {
	_, err := db.c.Exec("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	return err
}
