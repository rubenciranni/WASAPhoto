package database

func (db *appdbimpl) SetUserName(userID string, newUserName string) error {
	_, err := db.c.Exec("UPDATE User SET username = ? WHERE userID = ?", newUserName, userID)
	return err
}
