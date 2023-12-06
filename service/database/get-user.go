package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUser(userId string) (schema.User, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM User WHERE userId = ?", userId).Scan(&username)
	user := schema.User{
		UserId:   userId,
		Username: username,
	}
	return user, err
}
