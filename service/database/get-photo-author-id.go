package database

func (db *appdbimpl) GetPhotoAuthorId(photoId string) (string, error) {
	var authorId string
	err := db.c.QueryRow("SELECT authorId FROM Photo WHERE photoId = ?", photoId).Scan(&authorId)
	return authorId, err
}
