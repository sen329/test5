package main

import (
	"log"
	"net/http"

	controller "test5/Controller"
	"test5/Controller/admin"
	middleware "test5/Middleware"
	"test5/Routers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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
	route := router.PathPrefix("/api").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")
	route.HandleFunc("/getCurrentUserLogin", admin.GetCurrentUserLogin)
	route.HandleFunc("/getRolePermission", controller.GetRolePermission).Methods("GET")

	route.Handle("/", Routers.RouteIcon(route))
	route.Handle("/", Routers.RouteKsatriya(route))

	//	rune
	route.Handle("/", Routers.RouteRune(route))

	//	premium
	route.Handle("/", Routers.RoutePremium(route))

	//	energy
	route.Handle("/", Routers.RouteEnergy(route))

	//currency type
	route.Handle("/", Routers.RouteCurrency(route))

	//misc items
	route.Handle("/", Routers.RouteMiscItem(route))

	//box
	route.Handle("/", Routers.RouteBox(route))

	//chest
	route.Handle("/", Routers.RouteChest(route))

	// ---- Mail Subroute ---- //
	route.Handle("/", Routers.RouteMail(route))

	// ---- News Subroute ---- //
	route.Handle("/", Routers.RouteNews(route))

	// ---- Lotto Subroute ---- //
	route.Handle("/", Routers.RouteLotto(route))

	// ---- Gacha Subroute ---- //
	route.Handle("/", Routers.RouteGacha(route))

	// ---- Shop Subroute ---- //
	route.Handle("/", Routers.RouteLotus(route))
	route.Handle("/", Routers.RouteShop(route))

	// ---- Role Subroute ---- //
	route.Handle("/", Routers.RouteRole(route))

	route.Handle("/", Routers.RoutePlayers(route))

	route.Handle("/", Routers.RouteKsaRot(route))

	route.Handle("/", Routers.RouteUser(route))

	route.Handle("/", Routers.RouteMatches(route))

	route.Handle("/", Routers.RouteReports(route))

	route.Handle("/", Routers.RouteBlacklists(route))

	route.Handle("/", Routers.RouteVouchers(route))

	route.Handle("/", Routers.RouteWarning(route))

	route.Handle("/", Routers.RouteJudges(route))

	route.Handle("/", Routers.RoutePlayerStats(route))

	route.Handle("/", Routers.RouteGuild(route))

	route.Handle("/", Routers.RouteSeasons(route))

	route.Handle("/", Routers.RouteEvent(route))

	route.Handle("/", Routers.RouteMaintenance(route))

	route.Handle("/", Routers.RouteDailyReward(route))

	route.Handle("/", Routers.RouteBroadcast(route))

	route.Handle("/", Routers.RouteEventAnniversary(route))

	route.Handle("/", Routers.RouteClientVersion(route))

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
