package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type sendMessageRequest struct {
	// il nome delle robe json con maiuscola

	Text string `json:"messageText"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
	Foto string `json:"photoText"`
}

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //dichiarazioe funzione base
	parametroId := ps.ByName("id")
	parametroChat := ps.ByName("conversationId")

	idUser, err := strconv.Atoi(parametroId) //conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idChat, err := strconv.Atoi(parametroChat) //conversione da stringa a int
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

	var data sendMessageRequest
	if err := json.Unmarshal(body, &data); err != nil { //mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB converte da form a struttura
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("photo") // prende il file dalla form
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file) // legge il file
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// aggiungere modo per calcolare data, snippet e altri dati del messaggio
	err = rt.db.SendMessage(idUser, idChat, imageData, data.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
