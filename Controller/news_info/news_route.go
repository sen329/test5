package newsinfo

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(r *mux.Router) *mux.Router {

	route_news := r.PathPrefix("/news").Subrouter()
	route_news.Use(middleware.Middleware, middleware.CheckRoleMail)

	return r
}
