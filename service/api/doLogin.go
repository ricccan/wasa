package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type myData struct { // creo una classe per il json
	Nome string `json:"l_name"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	// Leggere il body
	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // gestisco gli errori
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	var data myData                                     // creo una variabile per il json
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	esiste, err := rt.db.SearchUser(data.Nome) // chiamo la funzione per cercare l'utente
	if err != nil {                            // gestisco gli errori
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if esiste == nil { // se l'utente non esiste chiamo la funzione che crea l'utente e lo inserisce nel database

		result, err := rt.db.DoLogin(data.Nome) // chiamo la funzione per fare il login creando l'utente
		if err != nil {                         // gestisco gli errori
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		str := strconv.Itoa(*result) // converto l'id in stringa

		response := map[string]string{ // creo un json con l'id
			"l_identifier": str,
		}

		strValue := response["l_identifier"] // prendo l'id

		intValue, err := strconv.Atoi(strValue) // converto l'id in int
		if err != nil {                         // gestisco gli errori
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, _ := generateToken(intValue) // genero il token

		// Imposta il Content-Type della risposta come JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Serializza i dati in formato JSON e scrivili nella risposta
		if err := json.NewEncoder(w).Encode(struct {
			Id    int
			Token string
		}{intValue, token}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

	} else { // se l'utente esiste genero il token

		str := strconv.Itoa(*esiste) // converto l'id in stringa

		response := map[string]string{ // creo un json con l'id
			"l_identifier": str,
		}

		strValue := response["l_identifier"] // prendo l'id

		intValue, err := strconv.Atoi(strValue) // converto l'id in int
		if err != nil {                         // gestisco gli errori
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, _ := generateToken(intValue) // genero il token
		// Imposta il Content-Type della risposta come JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Serializza i dati in formato JSON e scrivili nella risposta
		if err := json.NewEncoder(w).Encode(struct {
			Id    int
			Token string
		}{intValue, token}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

	}

}
