package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	tokenString := r.Header.Get("Authorization") // prendo il token
	if tokenString == "" {                       // gestisco se non c'è
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString) // lo verifico
	if err != nil {                 // gestisco se non è valido
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// prendo i parametri dall'url
	parametroId := ps.ByName("id")
	chatId := ps.ByName("conversationId")

	id, err := strconv.Atoi(parametroId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chat, err := strconv.Atoi(chatId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := rt.db.GetConversation(id, chat) // chiamo la funzione
	if err != nil {                                // gestione errori
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
