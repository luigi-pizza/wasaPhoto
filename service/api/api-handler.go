package api

import (
	"net/http"
)

// Ogni funzione in cui sei loggato deve checcare il bearer authorization
// userid non si deve passare perché è nell'header
// controlla se l'user è bannato o robe strane

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
