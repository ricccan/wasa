package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeFromGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //dichiarazioe funzione base
	parametroId := ps.ByName("id")
	gruppoId := ps.ByName("group")

	id, err := strconv.Atoi(parametroId) //conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gruppo, err := strconv.Atoi(gruppoId) //conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.RemoveFromGroup(id, gruppo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Imposta il Content-Type della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
