package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapingExecutionLog(t *testing.T) {

	router := mux.NewRouter()

	route := "/api/me/scraping_execution_log"
	router.HandleFunc(route, GetScrapingExecutionLog).Queries("limit", "{limit}").Methods("GET")

	request, _ := http.NewRequest("GET", route+"?limit=10&skip=0", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}
