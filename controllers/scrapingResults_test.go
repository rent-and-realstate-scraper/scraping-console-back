package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapedCities(t *testing.T) {

	router := mux.NewRouter()

	route := "/api/me/scraped_cities"
	router.HandleFunc(route, GetScrapedCities).Queries("scraping_id", "{scraping_id}").Methods("GET")

	request, _ := http.NewRequest("GET", route+"?scraping_id=scraping-airbnb-raspberryold--2019-3-5_15_09_28", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestGetScrapedResultsForCity(t *testing.T) {

	router := mux.NewRouter()

	route := "/api/me/scraped_results_for_city"
	router.HandleFunc(route, GetScrapedResultsForCity).Methods("GET")

	request, _ := http.NewRequest("GET", route+"?scraping_id=scraping-airbnb-raspberryold--2019-3-5_15_09_28&city_name=Alcobendas", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}
