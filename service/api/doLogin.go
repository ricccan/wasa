package api // dice che è delle api
import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type myData struct {
	// il nome delle robe json con maiuscola

	Nome string `json:"l_name"` // le parti del json che passiamo alla funzione tramite chiamata, si vedono dagli esempi nella pagina delle api grafiche
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	// Leggere il body
	body, err := io.ReadAll(r.Body) // readall funzione che legge il body
	if err != nil {                 // se ci sono errori
		http.Error(w, err.Error(), http.StatusBadRequest) // tira fuori errore
		return
	}
	defer r.Body.Close() // Chiudere il body dopo averlo letto

	// Deserializzare il JSON nel proprio tipo
	var data myData
	if err := json.Unmarshal(body, &data); err != nil { // mette il body in una variabile
		http.Error(w, err.Error(), http.StatusBadRequest) // se ci sono errori output
		return
	}

	esiste, err := rt.db.SearchUser(data.Nome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if esiste == nil {

		result, err := rt.db.DoLogin(data.Nome)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		str := strconv.Itoa(*result)

		response := map[string]string{
			"l_identifier": str,
		}

		strValue := response["l_identifier"]

		// Convert the string back to an integer
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, _ := generateToken(intValue)
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

	} else {
		str := strconv.Itoa(*esiste)

		response := map[string]string{
			"l_identifier": str,
		}

		strValue := response["l_identifier"]

		// Convert the string back to an integer
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, _ := generateToken(intValue)
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
