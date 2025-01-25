package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	rt.router.OPTIONS("/session", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Permette tutte le origini. Specifica un'origine particolare se necessario.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Metodi consentiti
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Header consentiti
		w.WriteHeader(http.StatusOK)                                                  // Risposta OK
	})
	// Register routes
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	rt.router.POST("/session", rt.doLogin) // funzione per il login, crea utente se non esiste o accede alle sue chat

	rt.router.POST("/users/:id/username", rt.setMyUsername) // funzione per cambiare il nome utente

	rt.router.POST("/users/:id/photo", rt.setMyPhoto) // funzione per cambiare la foto profilo

	rt.router.POST("/users/:id/groups", rt.addToGroup) // funzione per aggiungere un utente a un gruppo

	rt.router.DELETE("/users/:id/groups/:groups", rt.removeFromGroup) // funzione per rimuovere un utente da un gruppo

	rt.router.POST("/users/:id/conversations/:conversationId/groupName", rt.SetGroupName) // funzione per cambiare il nome del gruppo

	rt.router.POST("/users/:id/conversations/:conversationId/groupphoto", rt.SetGroupPhoto) // funzione per cambiare la foto del gruppo

	rt.router.POST("/users/:id/conversations/:conversationId/messages", rt.sendMessage) // funzione per inviare un messaggio

	rt.router.DELETE("/users/:id/conversations/:conversationId/messages/:messageId", rt.deleteMessage) // funzione per eliminare un messaggio

	rt.router.GET("/usernames/:name", rt.getUserId) // da aggiungere nelle api doc

	rt.router.DELETE("/users/:id/conversations/:conversationId/messages/:messageId/reactions", rt.uncommentMessage) // funzione per eliminare un commento (da cambiare)

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId/comments", rt.commentMessage) // funzione per commentare un messaggio

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId", rt.forwardMessage) // funzione per inoltrare un messaggio

	rt.router.GET("/users/:id/conversations", rt.getMyConversations) // funzione per ottenere le conversazioni dell'utente

	rt.router.GET("/users/:id/conversations/:conversationId/messages", rt.getMessages) // TODO: da aggiungere nelle api doc (funzione per ottenere i messaggi di una conversazione)

	rt.router.GET("/users/:id/photo", rt.getPhoto) // TODO: da aggiungere nelle api doc (funzione per ottenere la foto profilo avendo un id)

	rt.router.POST("/users/:id/conversations", rt.createGroup) // TODO: metterlo nelle api

	rt.router.POST("/users/:id", rt.createChat) // TODO: metterlo nelle api

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId/reactions", rt.reactMessage) // TODO AGGIUNGERE ALLE API

	return rt.router
}

// TODO: aggiungere funzione che calcola snippet
// TODO: aggiungere funzione che calcola checkmark
// TODO: aggiungere funzione che calcola icona
// TODO: aggiungere funzione che controlla il token
// TODO: cambiare uncomment facendo in modo che cambi solo la sezione respond e non l'intero messaggio
// TODO: sistemare sdoppiamento chat in caso di creazone con stesso utente
