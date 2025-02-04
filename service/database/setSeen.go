package database

func (db *appdbimpl) SetSeen(id int, conv int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "UPDATE visualizzato SET seen = 1 WHERE us = ? AND conv = ?" // query di aggiornamento

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(id, conv)
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
