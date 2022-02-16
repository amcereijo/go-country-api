package countries_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amcereijo/go-country-api/core/domain"
	"github.com/amcereijo/go-country-api/core/ports"
)

type HttpCountriesHandler struct {
	countriesService ports.CountriesService
}

func New(countriesService ports.CountriesService) *HttpCountriesHandler {
	return &HttpCountriesHandler{
		countriesService: countriesService,
	}
}

func (handler *HttpCountriesHandler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var country domain.Country
	json.Unmarshal(requestBody, &country)

	fmt.Printf("JSON country %+v\n", &country)

	var err error
	country, err = handler.countriesService.Create(
		country.Name,
		country.Capital,
	)

	if err != nil {
		fmt.Printf("Falied save country %+v\n", &err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Entity country %+v\n", &country)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(country)
}

func (handler *HttpCountriesHandler) GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := handler.countriesService.GetAll()

	if err != nil {
		fmt.Printf("Falied loading countries %+v\n", &err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(countries)
}
