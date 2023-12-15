/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rubenciranni/WASAPhoto/service/model/schema"
)

const (
	createUsersTable = `
		CREATE TABLE User (
			userID TEXT PRIMARY KEY,
			username TEXT NOT NULL
		);
	`
	createPhotoTable = `
		CREATE TABLE Photo (
			photoID TEXT PRIMARY KEY,
			authorId TEXT NOT NULL,
			caption TEXT,
			dateTime TEXT,
			FOREIGN KEY (authorId) REFERENCES User(userID)
		);
	`
	createCommentTable = `
		CREATE TABLE Comment (
			commentID TEXT PRIMARY KEY,
			photoID TEXT NOT NULL,
			authorId TEXT NOT NULL,
			text TEXT,
			dateTime TEXT,
			FOREIGN KEY (photoID) REFERENCES Photo(photoID),
			FOREIGN KEY (authorId) REFERENCES User(userID)
		);
	`
	createLikeTable = `
		CREATE TABLE Like (
			photoID TEXT NOT NULL,
			userID TEXT NOT NULL,
			PRIMARY KEY (photoID, userID),
			FOREIGN KEY (photoID) REFERENCES Photo(photoID),
			FOREIGN KEY (userID) REFERENCES User(userID)
		);
	`
	createFollowTable = `
		CREATE TABLE Follow (
			followerId TEXT NOT NULL,
			followedId TEXT NOT NULL,
			PRIMARY KEY (followerId, followedId),
			FOREIGN KEY (followerId) REFERENCES User(userID),
			FOREIGN KEY (followedId) REFERENCES User(userID)
		);
	`
	createBanTable = `
		CREATE TABLE Ban (
			bannerId TEXT NOT NULL,
			bannedId TEXT NOT NULL,
			PRIMARY KEY (bannerId, bannedId),
			FOREIGN KEY (bannerId) REFERENCES User(userID),
			FOREIGN KEY (bannedId) REFERENCES User(userID)
		);
	`
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	InsertUser(userID string, username string) error
	InsertPhoto(photoID string, authorId string, caption string, dateTime string) error
	InsertLike(photoID string, userID string) error
	InsertComment(commentID string, photoID string, authorId string, text string, dateTime string) error
	InsertFollow(followerId string, followedId string) error
	InsertBan(bannerId string, bannedId string) error

	DeletePhoto(photoID string) error
	DeleteLike(photoID string, userID string) error
	DeleteComment(commentID string) error
	DeleteFollow(followerId string, followedId string) error
	DeleteBan(bannerId string, bannedId string) error

	GetUserID(username string) (string, error)
	GetPhotosByUser(userID string, startDate string, startID string) ([]schema.Photo, error)
	GetPhotoAuthorId(photoID string) (string, error)
	GetLikes(photoID string, startID string) ([]schema.User, error)
	GetCommentAuthorId(commentID string) (string, error)
	GetComments(photoID string, startDate string, startID string) ([]schema.Comment, error)
	GetUsers(loggedInUserID string, username string, startID string) ([]schema.User, error)
	GetUser(userID string) (schema.User, error)
	GetFollowing(userID string, startID string) ([]schema.User, error)
	GetFollowers(userID string, startID string) ([]schema.User, error)
	GetUserProfile(userID string) (schema.UserProfile, error)
	GetStream(userID string, startDate string, startID string) ([]schema.Photo, error)

	SetUserName(userID string, newUserName string) error

	ExistsBan(bannerId string, bannedId string) (bool, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	var createStatements = map[string]string{
		"User":    createUsersTable,
		"Photo":   createPhotoTable,
		"Comment": createCommentTable,
		"Like":    createLikeTable,
		"Follow":  createFollowTable,
		"Ban":     createBanTable,
	}

	for tableName, sqlStatement := range createStatements {
		err := db.QueryRow("SELECT name FROM sqlite_master WHERE type = 'table' AND name = ?", tableName).Scan(&tableName)
		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.Exec(sqlStatement)
			if err != nil {
				return nil, fmt.Errorf("error creating database structure: %w", err)
			}
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
