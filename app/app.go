package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amcereijo/go-country-api/core/repositories/countries_repo"
	"github.com/amcereijo/go-country-api/core/services/countries_service"
	"github.com/amcereijo/go-country-api/database"
	"github.com/amcereijo/go-country-api/handler/countries_handler"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Intialize(envFileName string) {
	godotenv.Load(envFileName)

	config :=
		database.Config{
			ServerName: os.Getenv("MYSQL_HOST"),
			Port:       os.Getenv("MYSQL_PORT"),
			User:       os.Getenv("MYSQL_USER"),
			Password:   os.Getenv("MYSQL_PASS"),
			DB:         os.Getenv("MYSQL_DB_NAME"),
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	countriesRepository := countries_repo.New(database.Connector)
	countriesService := countries_service.New(countriesRepository)
	countriesHandler := countries_handler.New(countriesService)

	// create routes
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", countriesHandler.GetCountries).Methods("GET")
	postCreateRouter := router.Methods("POST").Subrouter()

	postCreateRouter.HandleFunc("/create", countriesHandler.CreateCountry)
	postCreateRouter.Use(countries_handler.Validator)

	a.Router = router
	a.DB = database.Connector
}

func (a *App) Run(envFileName string) {
	port := os.Getenv("PORT")
	log.Printf("Starting http server on port %s\n", port)
	log.Fatal((http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router)))
}
