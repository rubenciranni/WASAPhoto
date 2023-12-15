package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetComments(photoID string, startDate string, startId string) ([]schema.Comment, error) {
	var commentList []schema.Comment
	rows, err := db.c.Query(
		`
		SELECT Comment.commentId, Comment.authorId, User.username, Comment.text, Comment.dateTime
		FROM Comment JOIN User
		ON Comment.authorId = User.userID
		WHERE Comment.photoID = ? AND
		Comment.dateTime < ? OR (Comment.dateTime = ? AND Comment.commentId > ?)
		ORDER BY Comment.dateTime DESC, Comment.commentId ASC
		LIMIT 20
		`,
		photoID,
		startDate,
		startDate,
		startId,
	)
	if err != nil {
		return commentList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			commentId      string
			authorId       string
			authorUsername string
			text           string
			dateTime       string
		)
		if err := rows.Scan(&commentId, &authorId, &authorUsername, &text, &dateTime); err != nil {
			return commentList, err
		}
		commentList = append(commentList, schema.Comment{
			CommentID: commentId,
			Text:      text,
			Author:    schema.User{UserID: authorId, Username: authorUsername},
			DateTime:  dateTime,
		})
	}

	if err := rows.Err(); err != nil {
		return commentList, err
	}

	return commentList, err
}
