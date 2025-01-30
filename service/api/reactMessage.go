package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type React struct {
	Reazione string `json:"reaction"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) reactMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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
	userId := ps.ByName("id")
	messageId := ps.ByName("messageId")

	id, err := strconv.Atoi(userId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messid, err := strconv.Atoi(messageId) // conversione da stringa a int
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
	var data React
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}
	fmt.Println(data.Reazione) // output
	// chiamata funzione
	err = rt.db.ReactMessage(id, data.Reazione, messid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
