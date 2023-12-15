package database

func (db *appdbimpl) GetUserID(username string) (string, error) {
	var authorId string
	err := db.c.QueryRow("SELECT userID FROM User WHERE username = ?", username).Scan(&authorId)
	return authorId, err
}
