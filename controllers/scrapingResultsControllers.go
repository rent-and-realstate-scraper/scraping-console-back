package controllers

import (
	"net/http"
	"scraping-console-back/managers"
	"scraping-console-back/models"
	u "scraping-console-back/utils"

	"github.com/gorilla/mux"
)

var GetScrapedCities = func(w http.ResponseWriter, r *http.Request) {
	scrapedCities := models.ScrapedCities{}
	scrapingID := mux.Vars(r)["scraping_id"]
	data := scrapedCities.GetScrapedCities(scrapingID)
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

var GetScrapedResultsForCity = func(w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query()

	scrapingID := v.Get("scraping_id")
	cityName := v.Get("city_name")

	data := models.GetScrapingResultsForCity(cityName, scrapingID)

	geojson := managers.GenerateGeoJsonFromResult(data)
	var code int

	if data != nil {
		code = 200
	} else {
		code = 500
	}

	w.WriteHeader(code)

	var resp map[string]interface{} = make(map[string]interface{})
	resp["geojson"] = geojson
	u.Respond(w, resp)
}
