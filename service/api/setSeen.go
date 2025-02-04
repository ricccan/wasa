package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setSeen(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // dichiarazioe funzione base
	// gestione token

	// prendo i parametri
	userID := ps.ByName("id")
	conversationId := ps.ByName("conversationId")

	userId, err := strconv.Atoi(userID) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convid, err := strconv.Atoi(conversationId) // conversione da stringa a int
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// chiamata funzione
	err = rt.db.SetSeen(userId, convid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
