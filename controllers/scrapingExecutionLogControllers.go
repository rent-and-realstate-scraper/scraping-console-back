package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"scraping-console-back/models"
	u "scraping-console-back/utils"
	"strconv"
)

var GetScrapingExecutionLog = func(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(mux.Vars(r)["limit"])
	skip, _ := strconv.Atoi(mux.Vars(r)["skip"])
	offset, _ := mux.Vars(r)["offset"]

	var data = models.GetScrapingExecutionLog(limit, skip, offset)
	var code int
	code = 200

	w.WriteHeader(code)

	var resp map[string]interface{} = make(map[string]interface{})
	resp["data"] = data
	u.Respond(w, resp)
}
