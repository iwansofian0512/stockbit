package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/traveliodev/stockbit/2/handler"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/search-movies", handler.GetMoviesHander)

	recoveryHandler := negroni.NewRecovery()
	n := negroni.New(recoveryHandler)

	n.UseHandler(router)

	fmt.Println("Starting Travelio Channel Manager HTTP Server on 3001")
	http.ListenAndServe(":3001", n)
}

// type Route struct {
// 	Method  string
// 	Path    string
// 	Handler httprouter.Handle
// }

// func main() {
// 	recoveryHandler := negroni.NewRecovery()

// 	n := negroni.New(recoveryHandler)
// 	router := httprouter.New()
// 	SetRoute(router, GetRoute())
// 	n.UseHandler(router)

// }

// func GetRoute() []Route {
// 	return []Route{
// 		Route{"GET", "/api/health", handler.GetMoviesHander},
// 		// Route{"GET", "/activity-log/get", httpresponse.ResponseJson(handler.GetActivityLogHandler)},
// 	}
// }

// func SetRoute(router *httprouter.Router, route []Route) {
// 	for _, r := range route {
// 		router.Handle(r.Method, r.Path, r.Handler)
// 	}
// }
