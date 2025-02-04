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
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Ping() error
	DoLogin(username string) (*int, error)
	SetMyUsername(id int, username string) (string, error)
	SetMyPhoto(id int, photo []byte) (*User, error)
	AddToGroup(utente int, gruppo int) error
	RemoveFromGroup(utente int, gruppo int) error
	SetGroupName(id int, name string) error
	SetGroupPhoto(id int, photo []byte) error
	SendMessage(user int, chat int, photo []byte, text string) error
	DeleteMessage(utente int, gruppo int, messaggio int) error
	UncommentMessage(utente int, gruppo int) error
	CommentMessage(utente int, gruppo int, mess int, photo []byte, testo string) error
	ForwardMessage(id int, chat int, message int, invio int) error
	GetMyConversations(id int) (*[]Chat, error)
	GetConversation(id int, group int) (*Chat, error)
	GetMessages(id int, chat int) (*[]Message, error)
	SearchUser(username string) (*int, error)
	GetPhoto(id int) ([]byte, error)
	CreateGroup(nome string, utente int) error
	CreateChat(nome string, utente int) error
	AddToCollegamento(utente int, altro int, chat int64) error
	GetName(id int) (string, error)
	GetUserId(nome string) (*int, error)
	ReactMessage(id int, reazione string, mess int) error
	GetReactions(chat int) (*[]Reazione, error)
	SetSeen(id int, conv int) error
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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
