package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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

	conversationId := ps.ByName("id")
	id, err := strconv.Atoi(conversationId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // gestisco gli errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	var data setGroupNameRequest                        // dichiaro una variabile di tipo addToGroup
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	err = rt.db.CreateGroup(data.Nome, id) // chiamo la funzione AddToGroup con i dati passati
	if err != nil {                        // gestisco gli errori
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Imposta il Content-Type della risposta come JSON
	w.WriteHeader(http.StatusNoContent)

}
