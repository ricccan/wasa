package database

import "time"

func (db *appdbimpl) CreateChat(nome string, utente int) (int, error) { // creazione funzione, prende i parametri che ci servono
	// ritorna l'id dell'utente con un dato username
	query := "SELECT id FROM users WHERE username = ?" // dato il nome utente, ritorna il suo id

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return 0, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Query(nome)
	if err != nil {
		return 0, err
	}

	var altro *int // salvo l'id in una variabile
	for result.Next() {
		err = result.Scan(&altro)
		if err != nil {
			return 0, err
		}
	}

	// Controlla il numero di righe interessate
	err = result.Err()
	if err != nil {
		return 0, err
	}

	if altro == nil {
		return 0, err
	}

	// controllo se esiste già una conversazione tra i due utenti
	query = "SELECT conversations.id FROM conversations, us_con as primo, us_con as secondo WHERE conversations.id = primo.conv AND conversations.id = secondo.conv AND primo.us = ? AND secondo.us = ? AND conversations.grup = 'false'" // da cambiare

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return 0, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	controllo, err := stmt.Query(utente, altro)
	if err != nil {
		return 0, err
	}

	// Controlla il numero di righe interessate
	err = controllo.Err()
	if err != nil {
		return 0, err
	}

	// salvo il valore della query
	var prova *int
	for controllo.Next() {
		err = controllo.Scan(&prova)
		if err != nil {
			return 0, err
		}
	}

	// nel caso non ci siano già conversazioni
	if prova == nil {
		// creo la conversazione
		query = "INSERT INTO conversations (grup, lastchange) VALUES ('false', ?)"

		stmt, err = db.c.Prepare(query) // query
		if err != nil {
			return 0, err // se c è errore
		}
		defer stmt.Close() // Chiude lo statement preparato
		// Eseguire l'aggiornamento

		result2, err := stmt.Exec(time.Now().Unix())
		if err != nil {
			return 0, err
		}

		// Controlla il numero di righe interessate
		rowsAffected, err := result2.RowsAffected()
		if err != nil {
			return 0, err
		}
		if rowsAffected != 1 {
			return 0, err
		}

		chatid, err := result2.LastInsertId()
		if err != nil {
			return 0, err
		}

		// chiamo la funzione che crea il collegamento tra chat e utenti
		err = db.AddToCollegamento(utente, *altro, chatid)
		if err != nil {
			return 0, err
		}

		return int(chatid), nil
	} else {
		return *prova, nil
	}

}

func (db *appdbimpl) AddToCollegamento(utente int, altro int, chat int64) error { // creazione funzione, prende i parametri che ci servono
	// creo un collegamento tra la chat e il primo utente
	query := "INSERT INTO us_con (us, conv) VALUES (?,?)" // crea il collegamento

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(utente, chat)
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

	// creo il collegamento tra la chat e il secondo utente
	query = "INSERT INTO us_con (us, conv) VALUES (?,?)" // crea il secondo collegamento

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err = stmt.Exec(altro, chat)
	if err != nil {
		return err
	}

	// Controlla il numero di righe interessate
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return err
	}
	return nil
}
