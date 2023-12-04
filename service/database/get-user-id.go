package database

func (db *appdbimpl) GetUserId(username string) (string, error) {
	var authorId string
	err := db.c.QueryRow("SELECT userId FROM User WHERE username = ?", username).Scan(&authorId)
	return authorId, err
}
