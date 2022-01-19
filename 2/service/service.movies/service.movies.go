package movies

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	data "bitbucket.org/traveliodev/stockbit/2/service/service.movies/data"
)

type MoviesService struct {
	data data.MoviesData
}

func NewMoviesService() MoviesService {
	return MoviesService{
		data: data.MysqlMoviesData{},
	}
}

func (s MoviesService) GetMovies(payload Getmovies) (*GetMoviesResponse, error) {

	url := "http://www.omdbapi.com/?apikey=faf7e5bb&s=" + payload.Name + "&page=" + payload.Page

	LogId := GenerateID()

	var client = &http.Client{}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Connection", "close")

	if err != nil {
		return nil, err
	}

	timesTamp := time.Now()
	s.data.InsertLog(data.LogInsert{
		LogID:      LogId,
		Url:        url,
		Action:     "request",
		HttpMethod: "GET",
		Timestamp:  timesTamp,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	diff := time.Now().Sub(timesTamp)

	s.data.InsertLog(data.LogInsert{
		LogID:        LogId,
		Url:          url,
		Action:       "response",
		HttpMethod:   "GET",
		Response:     string(body),
		Timestamp:    time.Now(),
		ResponseTime: diff.Seconds(),
	})

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New(string(body))
	}

	if strings.Contains(string(body), "Error") {
		return nil, errors.New(string(body))
	}

	var result GetMoviesResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GenerateID() string {
	b := make([]byte, 16)

	_, err := rand.Read(b)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
