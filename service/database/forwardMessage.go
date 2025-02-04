package database

import (
	"time"
)

type Messaggio struct {
	Messag string
	Photo  []byte
}

func (db *appdbimpl) ForwardMessage(id int, chat int, message int, invio int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT messag, photo FROM messages WHERE conv = ? and id = ?" // prende il messaggio da inviare

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(chat, message)
	if err != nil {
		return err // se c è errore
	}
	defer rows.Close()

	var mess Messaggio
	for rows.Next() {
		err = rows.Scan(&mess.Messag, &mess.Photo)
		if err != nil {
			return err
		}
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return err
	}

	query = "INSERT INTO messages (conv, messag, photo, us, timestamp, forwarded) VALUES ( ?, ?, ?, ?, ?, true)" // copia il messaggio in un altra conversazione

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(invio, mess.Messag, mess.Photo, id, time.Now().Unix())
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

	risultato, err := stmt.Exec(time.Now().Unix(), mess.Messag, invio)
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
	fin, err := stmt.Exec(lastInsertId, invio, invio)
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
