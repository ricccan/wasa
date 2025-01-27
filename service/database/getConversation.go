package database

func (db *appdbimpl) GetConversation(id int, group int) (*Chat, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT id,grup_name,description,grup_photo,grup,lastchange,snippet   FROM conversations, us_con WHERE us_con.us = ? AND us_con.conv = conversations.id AND conversations.id = ? order by lastchange" // dato il nome utente, ritorna il suo id

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(id, group) // esegue la query
	if err != nil {
		return nil, err // se c è errore
	}
	defer rows.Close()

	var conv Chat
	for rows.Next() {
		err = rows.Scan(&conv.IdChat, &conv.GroupName, &conv.GroupDescription, &conv.GroupPhoto, &conv.Group, &conv.LastChange, &conv.Snippet)
		if err != nil {
			return nil, err
		}
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &conv, nil

}
