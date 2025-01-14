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
	rt.router.GET("/context", rt.wrap(rt.getContextReply)) // per ogni api devo fare un handler che richiama la funzione

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	rt.router.POST("/session", rt.doLogin) // chiamo l'endpoint mettendo percorso e nome

	rt.router.POST("/users/:id/username", rt.setMyUsername)

	rt.router.POST("/users/:id/photo", rt.setMyPhoto)

	rt.router.POST("/users/:id/groups", rt.addToGroup)

	rt.router.DELETE("/users/:id/groups/:group", rt.removeFromGroup)

	rt.router.POST("/users/:id/conversations/:conversationId/groupName", rt.SetGroupName)

	rt.router.POST("/users/:id/conversations/:conversationId/groupphoto", rt.SetGroupPhoto)

	rt.router.POST("/users/:id/conversations/:conversationId/messages", rt.sendMessage)

	rt.router.DELETE("/users/:id/conversations/:conversationId/messages/:messageId", rt.deleteMessage)

	rt.router.DELETE("/users/:id/conversations/:conversationId/messages/:messageId/comments", rt.uncommentMessage)

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId/comments", rt.commentMessage)

	rt.router.POST("/users/:id/conversations/:conversationId/messages/:messageId", rt.forwardMessage)

	rt.router.GET("/users/:id/conversations", rt.getMyConversations)

	rt.router.GET("/users/:id/conversations/:conversationId", rt.getConversation)

	rt.router.GET("/users/:id/conversations/:conversationId/messages", rt.getMessages) // prendo i messaggi di una conversazione (da aggiungere nelle api doc)

	rt.router.GET("/users/:id/photo", rt.getPhoto) // (da aggiungere nelle api doc)

	return rt.router
}
