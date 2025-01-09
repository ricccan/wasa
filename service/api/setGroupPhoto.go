package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SetGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	parametroId := ps.ByName("conversationId")

	id, err := strconv.Atoi(parametroId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	err = rt.db.SetGroupPhoto(id, imageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
