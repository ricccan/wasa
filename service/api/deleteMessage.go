package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	tokenString := r.Header.Get("Authorization") // prende il token dalla richiesta
	if tokenString == "" {                       // gestisco il caso in cui non ci sia il token
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString) // verifica il token
	if err != nil {                 // se il token non è valido
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// prendo i parametri dalla richiesta
	parametroId := ps.ByName("id")
	gruppoId := ps.ByName("conversationId")
	messaggio := ps.ByName("messageId")

	// conversione da stringa a int
	id, err := strconv.Atoi(parametroId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// conversione da stringa a int
	gruppo, err := strconv.Atoi(gruppoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// conversione da stringa a int
	mess, err := strconv.Atoi(messaggio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteMessage(id, gruppo, mess) // chiamo la funzione per eliminare il messaggio
	if err != nil {                             // gestisco gli errori
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Imposta il Content-Type della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
