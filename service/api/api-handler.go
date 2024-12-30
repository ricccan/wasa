package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/context", rt.wrap(rt.getContextReply)) // per ogni api devo fare un handler che richiama la funzione

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	rt.router.POST("/session", rt.doLogin) // chiamo l'endpoint mettendo percorso e nome

	rt.router.POST("/users/:id/username", rt.setMyUsername)

	rt.router.POST("/users/:id/photo", rt.setMyPhoto)
	return rt.router
}
