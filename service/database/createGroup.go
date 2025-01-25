package database

import "time"

func (db *appdbimpl) CreateGroup(nome string, utente int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "INSERT INTO conversations (grup, grup_name,lastchange) VALUES ('true',?, ?)" // creo un gruppo

	stmt, err := db.c.Prepare(query) // query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(nome, time.Now().Unix())
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

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	err = db.AddToGroup(utente, int(id))
	if err != nil {
		return err
	}

	return nil
}
