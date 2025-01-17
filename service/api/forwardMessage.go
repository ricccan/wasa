package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type forwardMessageRequest struct { // creazione di una struttura dati
	Chat int `json:"chatId"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	tokenString := r.Header.Get("Authorization") // prende il token dalla richiesta
	if tokenString == "" {                       // se il token è vuoto
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString) // verifica del token
	if err != nil {                 // se ci sono errori
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// prendo i parametri dall'url
	parametroId := ps.ByName("id")
	chatId := ps.ByName("conversationId")
	messageId := ps.ByName("messageId")

	// conversione da stringa a int
	id, err := strconv.Atoi(parametroId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// conversione da stringa a int
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// conversione da stringa a int
	message, err := strconv.Atoi(messageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // se ci sono errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	var data forwardMessageRequest                      // creo una variabile per contenere i dati
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile e gestisce errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: aggiungere funzione che calcola l'orario

	err = rt.db.ForwardMessage(id, chat, message, data.Chat) // chiamo la funzione forwardmessage
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
