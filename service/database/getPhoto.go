package database

func (db *appdbimpl) GetPhoto(id int) ([]byte, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT profile_picture FROM users WHERE id = ?" // query

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err // se c è errore
	}
	defer rows.Close()

	var foto []byte

	for rows.Next() {
		err = rows.Scan(&foto) // scan dei risultati
		if err != nil {
			return nil, err // se c è errore
		}
	}
	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foto, nil // ritorna la lista

}
