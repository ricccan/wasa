package database

func (db *appdbimpl) GetConversation(id int, group int) (*[]User, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT users.id,username,profile_picture FROM users,us_con WHERE users.id = us_con.us AND us_con.conv = ?" // dato il nome utente, ritorna il suo id

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(group) // esegue la query
	if err != nil {
		return nil, err // se c è errore
	}
	defer rows.Close()

	var usList []User
	var us User
	for rows.Next() {
		err := rows.Scan(&us.Id, &us.Username, &us.Image) // scansione dei risultati
		if err != nil {
			return nil, err // se c è errore
		}
		usList = append(usList, us)
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &usList, nil

}
