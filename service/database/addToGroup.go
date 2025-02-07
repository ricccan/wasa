package database

func (db *appdbimpl) AddToGroup(utente int, gruppo int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento

	/*
		la query seleziona l'id utente nel caso in cui fosse già
		all'interno del gruppo
	*/

	query := "SELECT us FROM us_con WHERE us = ? AND conv = ?" // query per vedere se l'utente è già nel gruppo

	stmt, err := db.c.Prepare(query) // prepara la query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	controllo, err := stmt.Query(utente, gruppo) // chiamo la query con i parametri
	if err != nil {
		return err
	}

	err = controllo.Err()
	if err != nil {
		return err
	}

	// salvo il valore della query
	var prova *int
	for controllo.Next() {
		err = controllo.Scan(&prova)
		if err != nil {
			return err
		}
	}

	// nel caso in cui l'utente non fosse già all'interno del gruppo
	if prova == nil {

		// collego l'utente con la conversazione. così facendo lo inserisco nel gruppo
		query = "INSERT INTO us_con (us, conv) VALUES (?,?)"

		stmt, err = db.c.Prepare(query) // prepara la query
		if err != nil {
			return err // se c è errore
		}
		defer stmt.Close() // Chiude lo statement preparato
		// Eseguire l'aggiornamento

		result, err := stmt.Exec(utente, gruppo) // chiamo la query con i parametri
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
