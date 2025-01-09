package database

func (db *appdbimpl) AddToGroup(utente int, gruppo int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "INSERT INTO us_con (us, conv) VALUES (?,?)" // da cambiare

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(utente, gruppo)
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
