package database

func (db *appdbimpl) GetName(id int) (string, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT username FROM users WHERE id = ?" // query

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return "", err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(id)
	if err != nil {
		return "", err // se c è errore
	}
	defer rows.Close()

	var nome string

	for rows.Next() {
		err = rows.Scan(&nome) // scan dei risultati
		if err != nil {
			return "", err // se c è errore
		}
	}
	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return "", err
	}

	return nome, nil // ritorna la lista

}
