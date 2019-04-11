package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"scraping-console-back/models"
	u "scraping-console-back/utils"
	"strconv"
)

var GetCount = func(w http.ResponseWriter, r *http.Request) {
	scrapingId, _ := strconv.ParseBool(mux.Vars(r)["scraped"])
	deviceId, _ := mux.Vars(r)["device_id"]

	var data = models.IndexGetScrapingCount(scrapingId, deviceId)
	var code int
	code = 200

	w.WriteHeader(code)

	var resp map[string]interface{} = make(map[string]interface{})
	resp["data"] = data
	u.Respond(w, resp)
}
