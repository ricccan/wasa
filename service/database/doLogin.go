package database

func (db *appdbimpl) DoLogin(username string) (*int, error) { // creazione funzione, prende i parametri che ci servono

	// Preparare la query
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

	return &id, nil

}
