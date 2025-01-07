package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type forwardMessageRequest struct {
	// il nome delle robe json con maiuscola

	Chat int `json:"chatId"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //dichiarazioe funzione base
	parametroId := ps.ByName("id")
	chatId := ps.ByName("conversationId")
	messageId := ps.ByName("messageId")

	id, err := strconv.Atoi(parametroId) //conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chat, err := strconv.Atoi(chatId) //conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := strconv.Atoi(messageId) //conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // se ci sono errori
		http.Error(w, err.Error(), http.StatusBadRequest) //tira fuori errore
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	var data forwardMessageRequest
	if err := json.Unmarshal(body, &data); err != nil { //mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	err = rt.db.ForwardMessage(id, chat, message, data.Chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
