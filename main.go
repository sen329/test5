package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
	"github.com/sen329/test5/Routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db := controller.Open()

	// defer db.Close()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router := mux.NewRouter()

	router.HandleFunc("/api/login", controller.Login).Methods("POST")

	// ---- normal route ---- //
	route := router.PathPrefix("/").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")

	route.Handle("/api", Routers.RouteIcon(route))
	route.Handle("/api", Routers.RouteKsatriya(route))

	//	rune
	route.Handle("/api", Routers.RouteRune(route))

	//	premium
	route.Handle("/api", Routers.RoutePremium(route))

	//	energy
	route.Handle("/api", Routers.RouteEnergy(route))

	//currency type
	route.Handle("/api", Routers.RouteCurrency(route))

	//box
	route.Handle("/api", Routers.RouteBox(route))

	//chest
	route.Handle("/api", Routers.RouteChest(route))

	// ---- Mail Subroute ---- //
	route.Handle("/api", Routers.RouteMail(route))

	// ---- Lotto Subroute ---- //
	route.Handle("/api", Routers.RouteLotto(route))

	// ---- Gacha Subroute ---- //
	route.Handle("/api", Routers.RouteGacha(route))

	// ---- Shop Subroute ---- //
	route.Handle("/api", Routers.RouteLotus(route))
	route.Handle("/api", Routers.RouteShop(route))

	// ---- Role Subroute ---- //
	route.Handle("/api", Routers.RouteRole(route))

	route.Handle("/api", Routers.RoutePlayers(route))

	route.Handle("/api", Routers.RouteKsaRot(route))

	route.Handle("/api", Routers.RouteUser(route))

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

// func mount(r *mux.Router, path string, handler http.Handler) {
// 	r.PathPrefix(path).Handler(
// 		http.StripPrefix(
// 			strings.TrimSuffix(path, "/"),
// 			handler,
// 		),
// 	)
// }
