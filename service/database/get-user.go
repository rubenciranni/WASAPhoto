package database

import "github.com/rubenciranni/WASAPhoto/service/model/schema"

func (db *appdbimpl) GetUser(userId string) (schema.User, error) {
	var user schema.User
	err := db.c.QueryRow("SELECT userId, username FROM User WHERE userId = ?", userId).Scan(&user)
	return user, err
}
