package routes

import (
	"github.com/gorilla/mux"
	"github.com/ozonebg/gofluence/context"
	"github.com/ozonebg/gofluence/middleware"
)

// NewRouter will return a prepared router will user defined routes.
func NewRouter(context *context.Context) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	routes := GetRoutes(context)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.Use(middleware.JwtAuthentication)

	return router
}
