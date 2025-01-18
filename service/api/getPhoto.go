package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	tokenString := r.Header.Get("Authorization") // controllo se il token è presente
	if tokenString == "" {                       // gestisco la sua assenza
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString) // controllo il token
	if err != nil {                 // in caso sia non valido
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// prendo i parametri
	parametroId := ps.ByName("id")

	id, err := strconv.Atoi(parametroId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := rt.db.GetPhoto(id) // chiamo la funzione
	if err != nil {                   // gestisco errori
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Imposta il Content-Type della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Serializza i dati in formato JSON e scrivili nella risposta
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
}
