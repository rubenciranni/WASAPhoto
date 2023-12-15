package database

func (db *appdbimpl) DeleteComment(commentID string) error {
	_, err := db.c.Exec("DELETE FROM Comment WHERE commentID = ?", commentID)
	return err
}
