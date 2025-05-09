package database

import (
	"time"
)

func (db *appdbimpl) CommentMessage(user int, chat int, mess int, photo []byte, text string) error { // creazione funzione, prende i parametri che ci servono

	// inserisco nel database il messaggio che è stato appena mandato
	query := "INSERT INTO messages (us, conv, risponde, photo, messag, timestamp) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(user, chat, mess, photo, text, time.Now().Unix())
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

	// avendo inserito il messaggio aggiorno lo snippet e la data della chat
	query = "UPDATE conversations SET lastchange = ?, snippet = ? WHERE id = ?" // query di aggiornamento

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	risultato, err := stmt.Exec(time.Now().Unix(), text, chat)
	if err != nil {
		return err
	}

	// Controlla il numero di righe interessate
	rowsAffected, err = risultato.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return err
	}

	// segna nel database che il messaggio appena mandato è stato visualizzato dall' utente che lo ha mandato
	query = "INSERT INTO visualizzato (us, mess, seen, conv) SELECT users.id, ?, 0, ? FROM users, us_con WHERE users.id = us_con.us AND us_con.conv = ?;"

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fin, err := stmt.Exec(lastInsertId, chat, chat)
	if err != nil {
		return err
	}

	// Controlla il numero di righe interessate
	rowsAffected, err = fin.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return err
	}

	err = db.SetSeen(user, chat)
	if err != nil {
		return err
	}

	return nil
}
