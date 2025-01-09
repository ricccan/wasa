package database

type Messaggio struct {
	snippet string
	messag  string
	photo   []byte
}

func (db *appdbimpl) ForwardMessage(id int, chat int, message int, invio int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT snippet, messag, photo FROM messages WHERE conv = ? and id = ?"

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	rows, err := stmt.Query(chat, message)
	if err != nil {
		return err // se c è errore
	}
	defer rows.Close()

	var mess Messaggio
	for rows.Next() {
		err = rows.Scan(&mess.snippet, &mess.messag, &mess.photo)
		if err != nil {
			return err
		}
	}

	query = "INSERT INTO messages (conv, snippet, messag, photo, us) VALUES (?, ?, ?, ?, ?)"

	stmt, err = db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(invio, mess.snippet, mess.messag, mess.photo, id)
	if err != nil {
		return err
	}

	// Controlla il numero di righe interessate
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return err
	}

	return nil
}
