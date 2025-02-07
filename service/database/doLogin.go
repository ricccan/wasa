package database

func (db *appdbimpl) SearchUser(username string) (*int, error) {
	// seleziono l'id dato l'username
	query := "SELECT id FROM users WHERE username = ?"

	stmt, err := db.c.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id *int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return id, nil
}

func (db *appdbimpl) DoLogin(username string) (*int, error) { // creazione funzione, prende i parametri che ci servono

	// creo un utente con l'username dato da input
	stmt, err := db.c.Prepare("INSERT INTO users (username) VALUES (?)") // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato

	// Esegui la query con valori specifici
	result, err := stmt.Exec(username) // esegue la query se non c è un return mettere underscore
	if err != nil {
		return nil, err
	}
	// Recupera l'ID della riga inserita
	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id := int(insertedID)

	// Controlla se la conversione è sicura (utile su sistemi a 32 bit)
	if int64(id) != insertedID {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected != 1 {
		return nil, err
	}

	return &id, nil

}
