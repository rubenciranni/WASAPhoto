package database

func (db *appdbimpl) GetCommentAuthorId(commentID string) (string, error) {
	var authorId string
	err := db.c.QueryRow("SELECT authorId FROM Comment WHERE commentID = ?", commentID).Scan(&authorId)
	return authorId, err
}
