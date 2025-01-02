package database

import "encoding/json"

type TextMessage struct {
	Id        int         `json:"t_id"`
	Messaggio string      `json:"t_textContent"`
	Preview   string      `json:"t_preview"`
	Sender    string      `json:"t_sender"`
	Data      json.Number `json:"t_date"` // da cambiare in data
	Checkmark int         `json:"t_checkmark"`
	Response  int         `json:"t_responseTo"`
	Photo     string      `json:"t_photoContent"` // da cambiare in formato foto
	Icon      string      `json:"t_icon"`         // formato foto
}
type Chat struct {
	// creo una classe per il json
	Id       int           `json:"c_id"`
	Utenti   []User        `json:"c_userIds"`
	Messaggi []TextMessage `json:"c_textMessageList"`
}

func (db *appdbimpl) AddToGroup(utente int, gruppo int) error { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "INSERT INTO us_con (us, conv) VALUES (?,?)" // da cambiare

	stmt, err := db.c.Prepare(query) //query
	if err != nil {
		return err // se c è errore
	}
	defer stmt.Close() // Chiude lo statement preparato
	// Eseguire l'aggiornamento

	result, err := stmt.Exec(utente, gruppo)
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
