package database

func (db *appdbimpl) ReactMessage(id int, reazione string, mess int) error { // creazione funzione, prende i parametri che ci servono

	// seleziona l'id di uno user avendo un suo messaggio e controlla se ha già fatto una reazione precedentemente a questa
	query := "SELECT us FROM reactions WHERE us = ? AND mess= ?" // query

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(id, mess)
	if err != nil {
		return err // se c è errore
	}
	defer rows.Close()

	var nome *int

	for rows.Next() {
		err = rows.Scan(&nome) // scan dei risultati
		if err != nil {
			return err // se c è errore
		}
	}
	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return err
	}

	// nel caso in cui non abbia messo reazioni precedenti
	if nome == nil {
		// creo una reazione
		query = "INSERT INTO reactions (us, mess, react) VALUES (?, ?, ?)" // query di inserimento

		stmt, err := db.c.Prepare(query) // query
		if err != nil {
			return err // se c è errore
		}
		defer stmt.Close() // Chiude lo statement preparato
		// Eseguire l'aggiornamento

		result, err := stmt.Exec(id, mess, reazione)
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

		// nel caso in cui abbia già messo una reazione precedentemente la modifico e la aggiorno a quella nuova
	} else {
		query = "UPDATE reactions SET react = ? WHERE us = ? AND mess= ?" // query di aggiornamento

		stmt, err = db.c.Prepare(query) // query
		if err != nil {
			return err // se c è errore
		}
		defer stmt.Close() // Chiude lo statement preparato
		// Eseguire l'aggiornamento

		result, err := stmt.Exec(reazione, id, mess)
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

	}

	return nil

}
