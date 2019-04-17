package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestScrapedCities(t *testing.T) {

	router := mux.NewRouter()

	route := "/api/scraping_results/scraped_cities"
	router.HandleFunc(route, GetScrapedCities).Queries("scraping_id", "{scraping_id}").Methods("GET")

	request, _ := http.NewRequest("GET", route+"?scraping_id=scraping-airbnb-raspberryold--2019-3-5_15_09_28", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestGetScrapedResultsForCity(t *testing.T) {

	router := mux.NewRouter()
	route := "/api/scraping_results/scraped_results_for_city"
	router.HandleFunc(route, GetScrapedResultsForCity).Methods("GET")

	request, _ := http.NewRequest("GET", route+"?scraping_id=scraping-airbnb-raspberryold--2019-3-5_15_09_28&city_name=Alcobendas", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.NotNil(t, response.Body, "intervals are filled")
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestProcessInfo(t *testing.T) {

	router := mux.NewRouter()
	route := "/api/scraping_results/process_info"
	router.HandleFunc(route, GetScrapedInfo).Methods("GET")

	request, _ := http.NewRequest("GET", route+"?scraping_id=scraping-airbnb-raspberryold--2019-3-5_15_09_28&device_id=raspberryold", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.NotNil(t, response.Body, "intervals are filled")
	assert.Equal(t, 200, response.Code, "OK response is expected")

}
