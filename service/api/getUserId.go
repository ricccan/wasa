package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type getId struct { // creo una classe per il json
	Nome string `json:"name"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) getUserId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
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

	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // gestisco gli errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	var data getId                                      // creo una variabile per il json
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, "csazzzo", http.StatusBadRequest) // se ci sono errori output
		return
	}
	// prendo i parametri dall'url

	result, err := rt.db.GetUserId(data.Nome) // chiamo la funzione
	if err != nil {                           // gestione errori
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
