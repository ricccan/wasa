package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct { // creo una classe per il json

	Id       int    `json:"u_id"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
	Username string `json:"u_username"`
	Image    string `json:"u_profileImage"`
}

type TextMessage struct { // creo una classe per il json
	Id        int         `json:"t_id"`
	Messaggio string      `json:"t_textContent"`
	Preview   string      `json:"t_preview"`
	Sender    string      `json:"t_sender"`
	Data      json.Number `json:"t_date"` // da cambiare in data
	Checkmark int         `json:"t_checkmark"`
	Response  int         `json:"t_responseTo"`
	Photo     string      `json:"t_photoContent"`
	Icon      string      `json:"t_icon"`
}

type Chat struct { // creo una classe per il json
	Id       int           `json:"c_id"`
	Utenti   []User        `json:"c_userIds"`         // lista di utenti
	Messaggi []TextMessage `json:"c_textMessageList"` // lista di messaggi
}

type addToGroup struct {
	Id_us    int `json:"adg_utente_id"`
	Id_group int `json:"adg_group_id"`
}

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	tokenString := r.Header.Get("Authorization") // prende il token dalla richiesta
	if tokenString == "" {                       // gestisco il caso in cui non ci sia il token
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString) // verifica il token
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized) // se il token non è valido
		return
	}

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // gestisco gli errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	var data addToGroup                                 // dichiaro una variabile di tipo addToGroup
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	err = rt.db.AddToGroup(data.Id_us, data.Id_group) // chiamo la funzione AddToGroup con i dati passati
	if err != nil {                                   // gestisco gli errori
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Imposta il Content-Type della risposta come JSON
	w.WriteHeader(http.StatusNoContent)

}
