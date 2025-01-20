package database

import (
	"time"
)

func (db *appdbimpl) CommentMessage(user int, chat int, mess int, photo []byte, text string) error { // creazione funzione, prende i parametri che ci servono
	// crea un messaggio nel database
	query := "INSERT INTO messages (us, conv, risponde, photo, messag, timestamp) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(user, chat, mess, photo, text, time.Now())
	if err != nil {
		return err
	}

	// Controlla il numero di righe interessate
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return err
	}

	return nil
}
