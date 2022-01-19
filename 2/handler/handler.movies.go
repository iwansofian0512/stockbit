package handler

import (
	"encoding/json"
	"net/http"

	movies "bitbucket.org/traveliodev/stockbit/2/service/service.movies"
)

func GetMoviesHander(writer http.ResponseWriter, request *http.Request) {

	queryValues := request.URL.Query()
	name := queryValues.Get("searchWord")
	page := queryValues.Get("pagination")

	s := movies.NewMoviesService()
	results, err := s.GetMovies(movies.Getmovies{Name: name, Page: page})

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}

	json.NewEncoder(writer).Encode(results)
	return
}
