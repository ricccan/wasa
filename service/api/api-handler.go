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

	rt.router.POST("/session", rt.doLogin) // funzione per il login, crea utente se non esiste o accede alle sue chat V V

	rt.router.POST("/users/:id/username", rt.setMyUsername) // funzione per cambiare il nome utente V V

	rt.router.POST("/users/:id/photo", rt.setMyPhoto) // funzione per cambiare la foto profilo V V

	rt.router.POST("/users/:id/groups", rt.addToGroup) // funzione per aggiungere un utente a un gruppo V V

	rt.router.DELETE("/users/:id/groups/:groups", rt.removeFromGroup) // funzione per rimuovere un utente da un gruppo V V

	rt.router.POST("/users/:id/conversations/:conversationId/groupName", rt.SetGroupName) // funzione per cambiare il nome del gruppo V V

	rt.router.POST("/users/:id/conversations/:conversationId/groupphoto", rt.SetGroupPhoto) // funzione per cambiare la foto del gruppo V V

	rt.router.POST("/users/:id/conversations/:conversationId/messages", rt.sendMessage) // funzione per inviare un messaggio V V

	rt.router.DELETE("/users/:id/conversations/:conversationId/messages/:messageId", rt.deleteMessage) // funzione per eliminare un messaggio V V

	rt.router.GET("/usernames/:name", rt.getUserId) // funzione per ottenere l'id di un utente dato il nome V V

	rt.router.DELETE("/users/:id/conversations/:conversationId/messages/:messageId/reactions", rt.uncommentMessage) // funzione per eliminare una reaction (da cambiare) V V

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId/comments", rt.commentMessage) // funzione per commentare un messaggio V V

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId", rt.forwardMessage) // funzione per inoltrare un messaggio V V

	rt.router.GET("/users/:id/conversations", rt.getMyConversations) // funzione per ottenere le conversazioni dell'utente V V

	rt.router.GET("/users/:id/conversations/:conversationId", rt.getConversation) // V

	rt.router.GET("/users/:id/conversations/:conversationId/messages", rt.getMessages) // funzione per ottenere i messaggi di una conversazione V V

	rt.router.GET("/users/:id/photo", rt.getPhoto) // funzione per ottenere la foto profilo avendo un id V V

	rt.router.POST("/users/:id/conversations", rt.createGroup) // funzione che crea un gruppo V V

	rt.router.POST("/users/:id", rt.createChat) // funzione che crea una chat diretta tra user V V

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId/reactions", rt.reactMessage) // funzione per reagire ad un messaggio V V

	rt.router.GET("/users/:id/conversations/:conversationId/messages/:messageId/reactions", rt.getReactions) // funzione per ottenere le reazioni ad un messaggio V V

	rt.router.POST("/users/:id/conversations/:conversationId", rt.setSeen) // funzione per segnare come visualizzato un messaggio da un utente TODO: da aggiungere alle api V V

	return rt.router
}

// TODO: aggiungere funzione che calcola checkmark
