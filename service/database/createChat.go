package database

func (db *appdbimpl) CreateChat(nome string, utente int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT id FROM users WHERE username = ?" // da cambiare

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Query(nome)
	if err != nil {
		return err
	}

	var altro int
	for result.Next() {
		err = result.Scan(&altro)
		if err != nil {
			return err
		}
	}

	// Controlla il numero di righe interessate
	err = result.Err()
	if err != nil {
		return err
	}

	query = "SELECT conversations.id FROM conversations, us_con as primo, us_con as secondo WHERE conversations.id = primo.conv AND conversations.id = secondo.conv AND primo.us = ? AND secondo.us = ?" // da cambiare

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	controllo, err := stmt.Query(utente, altro)
	if err != nil {
		return err
	}

	// Controlla il numero di righe interessate
	err = controllo.Err()
	if err != nil {
		return err
	}

	var prova *int
	for controllo.Next() {
		err = controllo.Scan(&prova)
		if err != nil {
			return err
		}
	}

	if prova == nil {
		query = "INSERT INTO conversations (grup) VALUES ('false')"

		stmt, err = db.c.Prepare(query) // query
		if err != nil {
			return err // se c è errore
		}
		defer stmt.Close() // Chiude lo statement preparato
		// Eseguire l'aggiornamento

		result2, err := stmt.Exec()
		if err != nil {
			return err
		}

		// Controlla il numero di righe interessate
		rowsAffected, err := result2.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected != 1 {
			return err
		}

		chatid, err := result2.LastInsertId()
		if err != nil {
			return err
		}

		db.AddToCollegamento(utente, altro, chatid)

		return nil
	}
	return nil

}

func (db *appdbimpl) AddToCollegamento(utente int, altro int, chat int64) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "INSERT INTO us_con (us, conv) VALUES (?,?)" // da cambiare

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

	query = "INSERT INTO us_con (us, conv) VALUES (?,?)" // da cambiare

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
