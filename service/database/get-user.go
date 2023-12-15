package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUser(userID string) (schema.User, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM User WHERE userID = ?", userID).Scan(&username)
	user := schema.User{
		UserID:   userID,
		Username: username,
	}
	return user, err
}
