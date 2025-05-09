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
	GetConversation(id int, group int) (*[]User, error)
	GetMessages(id int, chat int) (*[]Message, error)
	SearchUser(username string) (*int, error) // 1
	GetPhoto(id int) ([]byte, error)
	CreateGroup(nome string, utente int) error
	CreateChat(nome string, utente int) (int, error)
	AddToCollegamento(utente int, altro int, chat int64) error // 2
	GetName(id int) (string, error)
	GetUserId(nome string) (*int, error)
	ReactMessage(id int, reazione string, mess int) error
	GetReactions(chat int) (*[]Reazione, error)
	SetSeen(id int, conv int) error // 3
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

	tables := map[string]string{
		"users": "CREATE TABLE users(id integer PRIMARY KEY AUTOINCREMENT, username text UNIQUE, profile_picture BLOB)",

		"reactions": `CREATE TABLE reactions( 
		 us integer,
		 mess integer,
		 react text,
		 primary key (us,mess),
		foreign key (mess) references messages(id),
		 foreign key (us)	references users(id)    
		 )`,

		"conversations": "CREATE TABLE conversations(id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, grup BOOLEAN NOT NULL, description TEXT, grup_photo BLOB, grup_name TEXT, lastchange INTEGER, snippet TEXT)",

		"messages": `CREATE TABLE messages(id         INTEGER  PRIMARY KEY AUTOINCREMENT,
		timestamp INTEGER,
		messag     TEXT,
		checkmarks INTEGER  CHECK (checkmarks < 3),
		conv       INTEGER,
		us         INTEGER,
		risponde   INTEGER, photo BLOB, forwarded BOOLEAN,
		FOREIGN KEY (
			conv
		)
		REFERENCES conversations (id),
		FOREIGN KEY (
			us
		)
		REFERENCES users (id),
		FOREIGN KEY (
			risponde
		)
		REFERENCES messages (id),
		check ( id <> risponde)
		)`,

		"us_con": `CREATE TABLE us_con( 
        us integer,
        conv integer,
        primary key (us,conv),
		foreign key (conv) references conversations(id),
        foreign key (us)	references users(id)    
		)`,

		"visualizzato": `CREATE TABLE visualizzato( 
         us integer,
         mess integer,
		 seen integer, "conv" INTEGER REFERENCES "conversations"("id"),
         primary key (us,mess),
	    foreign key (mess) references messages(id),
         foreign key (us)	references users(id)    
		)`,
	}

	// TODO da sistemare (roba di alex)
	for table, queryTable := range tables {
		var tableName string
		err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name=?;`, table).Scan(&tableName)
		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.Exec(queryTable)
			if err != nil {
				return nil, fmt.Errorf("error creating database structure: %w", err)
			}
		}
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	/* var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	*/
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
