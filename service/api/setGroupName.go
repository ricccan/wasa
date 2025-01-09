package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type setGroupNameRequest struct {
	// il nome delle robe json con maiuscola
	Nome string `json:"sgn_name"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) SetGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	conversationId := ps.ByName("conversationId")

	id, err := strconv.Atoi(conversationId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // se ci sono errori
		http.Error(w, err.Error(), http.StatusBadRequest) // tira fuori errore
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	// Deserializzare il JSON nel proprio tipo
	var data setGroupNameRequest
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	err = rt.db.SetGroupName(id, data.Nome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
