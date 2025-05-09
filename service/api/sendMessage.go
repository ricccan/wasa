package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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

	// prendo i parametri
	parametroId := ps.ByName("id")
	parametroChat := ps.ByName("conversationId")

	idUser, err := strconv.Atoi(parametroId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idChat, err := strconv.Atoi(parametroChat) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isPhoto := r.FormValue("IsPhoto")              // prende il valore dalla form
	messageType, err := strconv.ParseBool(isPhoto) // converte il valore in booleano

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB converte da form a struttura
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var imageData []byte
	file, _, err := r.FormFile("photo") // prende il file dalla form
	if err != nil && messageType {      // nel caso non sia una foto
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err == nil { // nel caso sia una foto
		defer file.Close()
		imageData, err = io.ReadAll(file) // legge il file
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	mess := r.FormValue("messageText") // prende il file dalla form

	// aggiungere modo per calcolare data, snippet e altri dati del messaggio
	err = rt.db.SendMessage(idUser, idChat, imageData, mess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
