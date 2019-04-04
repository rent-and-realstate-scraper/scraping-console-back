package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"scraping-console-back/models"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/user/new", CreateAccount)
	body := models.Account{Email:"email@email", Password:"password"}
	json, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", "/api/user/new", bytes.NewBuffer(json))
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)

}
