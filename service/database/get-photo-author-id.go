package database

func (db *appdbimpl) GetPhotoAuthorId(photoID string) (string, error) {
	var authorId string
	err := db.c.QueryRow("SELECT authorId FROM Photo WHERE photoID = ?", photoID).Scan(&authorId)
	return authorId, err
}
