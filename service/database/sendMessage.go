package database

import (
	"time"
)

func (db *appdbimpl) SendMessage(user int, chat int, photo []byte, text string) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento

	query := "INSERT INTO messages (us, conv, photo, messag, timestamp, forwarded) VALUES (?, ?, ?, ?, ?, false)"

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(user, chat, photo, text, time.Now().Unix())
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

	return nil
}
