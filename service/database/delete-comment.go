package database

func (db *appdbimpl) DeleteComment(commentId string) error {
	_, err := db.c.Exec("DELETE FROM Comment WHERE commentId = ?", commentId)
	return err
}
