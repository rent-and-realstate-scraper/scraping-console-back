package controllers

import (
	"net/http"
	"scraping-console-back/models"
	u "scraping-console-back/utils"
	"strconv"
)

var GetScrapingExecutionLog = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	limit, _ := strconv.Atoi(v.Get("limit"))
	skip, _ := strconv.Atoi(v.Get("skip"))
	offset := v.Get("order")

	var data = models.GetScrapingExecutionLog(limit, skip, offset)
	var code int
	code = 200

	w.WriteHeader(code)

	var resp map[string]interface{} = make(map[string]interface{})
	resp["data"] = data
	u.Respond(w, resp)
}
