package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	// il nome delle robe json con maiuscola

	Id       int    `json:"u_id"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
	Username string `json:"u_username"`
	Image    string `json:"u_profileImage"`
	Bio      string `json:"u_bio"`
	Name     string `json:"u_firstName"`
	Lastname string `json:"u_lastName"`
}

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

type addToGroup struct {
	Id_us    int `json:"adg_utente_id"`
	Id_group int `json:"adg_group_id"`
}

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // se ci sono errori
		http.Error(w, err.Error(), http.StatusBadRequest) // tira fuori errore
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	// Deserializzare il JSON nel proprio tipo
	var data addToGroup
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	err = rt.db.AddToGroup(data.Id_us, data.Id_group) // da creare la funzione e mettere i dati che passo alla query
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Imposta il Content-Type della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
