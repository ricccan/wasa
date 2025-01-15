package database

type User struct {
	// il nome delle robe json con maiuscola

	Id       int     `json:"u_id"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
	Username string  `json:"u_username"`
	Image    *string `json:"u_profileImage"`
}

func (db *appdbimpl) SetMyPhoto(id int, photo []byte) (*User, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "UPDATE users SET profile_picture = ? WHERE id = ?"

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return nil, err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(photo, id)
	if err != nil {
		return nil, err
	}

	// Controlla il numero di righe interessate
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected != 1 {
		return nil, err
	}

	query = "SELECT * FROM users WHERE id = ?"

	stmt, err = db.c.Prepare(query) // query
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

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Image)
		if err != nil {
			return nil, err
		}
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil

}
