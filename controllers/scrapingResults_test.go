package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapingResults(t *testing.T) {

	router := mux.NewRouter()

	route := "/api/me/scraped_cities"
	router.HandleFunc(route, GetScrapedCIties).Queries("scraping_id", "{scraping_id}").Methods("GET")

	request, _ := http.NewRequest("GET", route+"?scraping_id=scraping-fotocasa-raspberryWk--2019-4-5_18_06_58", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response is expected")



}
