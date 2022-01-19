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
