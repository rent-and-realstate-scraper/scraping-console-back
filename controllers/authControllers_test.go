package controllers

import (
	"bytes"
	"encoding/json"
	"scraping-console-back/models"
	"testing"
		"github.com/stretchr/testify/assert"
		"net/http/httptest"
		"net/http"

)

func TestCreateAccount(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/user/new", CreateAccount)
	body := models.Account{Email:"111", Password:"222"}
	json, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", "/api/user/new", bytes.NewBuffer(json))
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}
