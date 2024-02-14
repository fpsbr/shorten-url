package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) shortenURL(w http.ResponseWriter, r *http.Request) {
	var shortenInput struct {
		Url string `json:"url"`
	}

	err := app.readJSON(w, r, &shortenInput)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	newUrl, err := app.shortener.ShortenURL(shortenInput.Url)
	if err != nil {
		app.serveErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"url": newUrl}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}
}

func (app *application) showShortenURL(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	key := params.ByName("key")

	longUrl, err := app.shortener.GetUrl(key)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"url": longUrl}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}
}
