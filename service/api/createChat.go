package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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

	userId := ps.ByName("id")
	id, err := strconv.Atoi(userId) // conversione da stringa a int
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

	result, err := rt.db.CreateChat(data.Nome, id) // chiamo la funzione AddToGroup con i dati passati
	if err != nil {                                // gestisco gli errori
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
