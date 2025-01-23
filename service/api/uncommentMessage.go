package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	// gestione token
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// gestione parametri
	parametroId := ps.ByName("id")

	messaggio := ps.ByName("messageId")

	id, err := strconv.Atoi(parametroId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mess, err := strconv.Atoi(messaggio) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// chiamata funzione
	err = rt.db.UncommentMessage(id, mess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Imposta il Content-Type della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
