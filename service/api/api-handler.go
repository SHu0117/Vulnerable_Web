package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {

	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
