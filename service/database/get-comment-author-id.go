package database

func (db *appdbimpl) GetCommentAuthorId(commentId string) (string, error) {
	var authorId string
	err := db.c.QueryRow("SELECT authorId FROM Comment WHERE commentId = ?", commentId).Scan(&authorId)
	return authorId, err
}
