package database

func (db *appdbimpl) DeleteMessage(utente int, gruppo int, messaggio int) error { // creazione funzione, prende i parametri che ci servono
	// cancella il messaggio con l'id passato dai parametri della funzione
	query := "DELETE FROM `messages` WHERE (`us` = ? AND `conv` = ? AND `id` = ?)" // cancella un messaggio

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(utente, gruppo, messaggio)
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
