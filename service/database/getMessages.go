package database

type Message struct {
	// il nome delle robe json con maiuscola

	IdMess     int
	Timestamp  *int
	Messaggio  *string
	Checkmarks *int
	User       int
	Username   string
	Risposta   *int
	Foto       []byte
	Inoltrato  *bool
}

func (db *appdbimpl) GetMessages(id int, gruppo int) (*[]Message, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT messages.id,timestamp,messag,checkmarks,us,username,risponde,photo,forwarded FROM messages,users WHERE conv = ? and us = users.id" // query

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(gruppo)
	if err != nil {
		return nil, err // se c è errore
	}
	defer rows.Close()

	var listaMess []Message
	var mess Message

	for rows.Next() {
		err = rows.Scan(&mess.IdMess, &mess.Timestamp, &mess.Messaggio, &mess.Checkmarks, &mess.User, &mess.Username, &mess.Risposta, &mess.Foto, &mess.Inoltrato) // scan dei risultati
		if err != nil {
			return nil, err // se c è errore
		}
		listaMess = append(listaMess, mess) // appende alla lista
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &listaMess, nil // ritorna la lista

}

func (db *appdbimpl) GetOneMessage(id int) (*[]Message, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT messages.id,timestamp,messag,checkmarks FROM messages WHERE messages.id = ?" // query

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

	var listaMess []Message // forse da cambiare struct perchè non ci sono tutti i campi
	var mess Message

	for rows.Next() {
		err = rows.Scan(&mess.IdMess, &mess.Timestamp, &mess.Messaggio, &mess.Checkmarks) // scan dei risultati
		if err != nil {
			return nil, err // se c è errore
		}
		listaMess = append(listaMess, mess) // appende alla lista
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &listaMess, nil // ritorna la lista

}
