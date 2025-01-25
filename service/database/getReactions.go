package database

type Reazione struct {
	Reazione string
	User     string
}

func (db *appdbimpl) GetReactions(chat int) (*[]Reazione, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT react, users.username FROM reactions,users WHERE mess = ? and us = users.id" // query

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(chat)
	if err != nil {
		return nil, err // se c è errore
	}
	defer rows.Close()

	var listaReaz []Reazione
	var reaz Reazione
	for rows.Next() {
		err = rows.Scan(&reaz.Reazione, &reaz.User)
		if err != nil {
			return nil, err
		}
		listaReaz = append(listaReaz, reaz)
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &listaReaz, nil
} // ritorna la lista
