package database

func (db *appdbimpl) SetGroupName(id int, name string) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "UPDATE conversations SET grup_name = ? WHERE id = ?"

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(name, id)
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
