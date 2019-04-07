package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"scraping-console-back/models"
	u "scraping-console-back/utils"
)

var GetScrapedCities = func(w http.ResponseWriter, r *http.Request) {
	scrapedCities := models.ScrapedCities{}
	scrapingId:= mux.Vars(r)["scraping_id"]
	data:= scrapedCities.GetScrapedCities(scrapingId)
	var code int
	if data != nil && data.ScrapedCities != nil {
		code = 200
	} else {
		code = 500
	}

	w.WriteHeader(code)

	var resp map[string]interface{} = make(map[string]interface{})
	resp["data"] = data
	u.Respond(w, resp)
}

