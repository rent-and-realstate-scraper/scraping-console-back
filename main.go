package main

import (
	"log"
	"net/http"
	"os"
	"scraping-console-back/controllers"
	"scraping-console-back/middlewares"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/auth/sign_in", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/auth/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/scraping_results/scraped_cities", controllers.GetContactsFor).Queries("scraping_id", "{scraping_id}").Methods("GET")
	router.HandleFunc("/api/scraping_results/scraping_execution_log", controllers.GetScrapingExecutionLog).Methods("GET")
	router.HandleFunc("/api/scraping_results/scraped_results_for_city", controllers.GetScrapedResultsForCity).Methods("GET")
	router.HandleFunc("/api/scraping_results/process_info", controllers.GetScrapedInfo).Methods("GET")

	//router.Use(middlewares.JwtAuthentication) //attach JWT auth middleware
	router.Use(middlewares.MiddlewareLogger) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
