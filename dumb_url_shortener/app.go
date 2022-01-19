package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const maxCapacity = 10

type App struct {
	Router *mux.Router
	Mapper *UrlMapper
}

type RequestBody struct {
	OriginalUrl string `json:"original_url"`
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Mapper = NewUrlMapper(NewSequenceGenerator(), maxCapacity)
	a.initRoutes()
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) createShortUrl(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	host := r.Host
	token, err := a.Mapper.GenerateShortUrlToken(requestBody.OriginalUrl)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, fmt.Sprintf("%s/%s", host, token))
}

func (a *App) getOriginalUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	originalUrl := a.Mapper.GetUrlByShortUrl(vars["token"])
	if originalUrl == "" {
		respondWithError(w, http.StatusNotFound, "Unrecognized URL")
		return
	}

	respondWithJSON(w, http.StatusOK, originalUrl)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/url", a.createShortUrl).Methods("POST")
	a.Router.HandleFunc("/{token}", a.getOriginalUrl).Methods("GET")
}
