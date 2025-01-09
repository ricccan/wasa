package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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

	isPhoto := r.FormValue("IsPhoto") // prende il file dalla form
	messageType, err := strconv.ParseBool(isPhoto)

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
	if err != nil && messageType {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err == nil {
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
