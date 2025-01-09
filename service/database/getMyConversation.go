package database

type Chat struct {
	// il nome delle robe json con maiuscola

	IdChat           int
	GroupName        *string
	GroupDescription *string
	GroupPhoto       []byte
	Group            bool
}

func (db *appdbimpl) GetMyConversations(id int) (*[]Chat, error) { // creazione funzione, prende i parametri che ci servono
	// Query di aggiornamento
	query := "SELECT id,grup_name,description,grup_photo,grup   FROM conversations, us_con WHERE us_con.us = ? AND us_con.conv = conversations.id"

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

	var listaChat []Chat
	var chat Chat
	for rows.Next() {
		err = rows.Scan(&chat.IdChat, &chat.GroupName, &chat.GroupDescription, &chat.GroupPhoto, &chat.Group)
		if err != nil {
			return nil, err
		}
		listaChat = append(listaChat, chat)
	}

	return &listaChat, nil

}
