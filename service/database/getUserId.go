package database

func (db *appdbimpl) GetUserId(nome string) (*int, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT id FROM users WHERE username = ?" // query

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(nome)
	if err != nil {
		return nil, err // se c è errore
	}
	defer rows.Close()

	var id *int

	for rows.Next() {
		err = rows.Scan(&id) // scan dei risultati
		if err != nil {
			return nil, err // se c è errore
		}
	}
	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return id, nil // ritorna la lista

}
