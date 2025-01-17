package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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
	parametroChat := ps.ByName("conversationId")
	parametroMessage := ps.ByName("messageId")

	// conversione da stringhe a int
	idUser, err := strconv.Atoi(parametroId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// conversione da stringhe a int
	idChat, err := strconv.Atoi(parametroChat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// conversione da stringhe a int
	idMessage, err := strconv.Atoi(parametroMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isPhoto := r.FormValue("IsPhoto")              // prende il valore del form, true se è una foto, false se non lo è
	messageType, err := strconv.ParseBool(isPhoto) // conversione da stringhe a bool per sapere se c'è un'immagine

	if err != nil { // gestisco gli errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB converte da form a struttura
	if err != nil {                      // gestisco gli errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var imageData []byte                // dichiaro un array di byte per l'immagine
	file, _, err := r.FormFile("photo") // prende il file dalla form con l'immagine
	if err != nil && messageType {      // gestisco gli errori e controllo se è una foto o no
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err == nil { // se non ci sono errori
		defer file.Close()
		imageData, err = io.ReadAll(file) // legge il file
		if err != nil {                   // gestisco gli errori
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	mess := r.FormValue("messageText") // prende il file dalla form

	// TODO: aggiungere funzione per calcolare data e snippet

	err = rt.db.CommentMessage(idUser, idChat, idMessage, imageData, mess) // chiamo la funzione commentmessage passando i parametri
	if err != nil {                                                        // gestisco gli errori
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
