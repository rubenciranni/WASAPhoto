package database

func (db *appdbimpl) SetUserName(userId string, newUserName string) error {
	_, err := db.c.Exec("UPDATE User SET username = ? WHERE userId = ?", newUserName, userId)
	return err
}
