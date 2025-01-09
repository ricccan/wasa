package database

type Message struct {
	// il nome delle robe json con maiuscola

	IdMess     int
	Snippet    *string
	Timestamp  *int
	Messaggio  *string
	Checkmarks *int
	User       int
	Risposta   *int
	Foto       []byte
}

func (db *appdbimpl) GetMessages(id int, gruppo int) (*[]Message, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT id,snippet,timestamp,messag,checkmarks,us,risponde,photo FROM messages WHERE conv = ?" // query

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
		err = rows.Scan(&mess.IdMess, &mess.Snippet, &mess.Timestamp, &mess.Messaggio, &mess.Checkmarks, &mess.User, &mess.Risposta, &mess.Foto) // scan dei risultati
		if err != nil {
			return nil, err // se c è errore
		}
		listaMess = append(listaMess, mess) // appende alla lista
	}

	return &listaMess, nil // ritorna la lista

}
