package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"scraping-console-back/models"
	u "scraping-console-back/utils"
)

var GetScrapedCIties = func(w http.ResponseWriter, r *http.Request) {
	scrapedCities := models.ScrapedCities{}
	scrapingId:= mux.Vars(r)["scraping_id"]
	data, code := scrapedCities.GetScrapedCities(scrapingId)
	if (code != 200) {
		w.WriteHeader(code)
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

