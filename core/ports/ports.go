package ports

import "github.com/amcereijo/go-country-api/core/domain"

type CountriesRepository interface {
	SaveCountry(country domain.Country) (domain.Country, error)
	GetAll() ([]domain.Country, error)
}

type CountriesService interface {
	Create(name string, capital string) (domain.Country, error)
	GetAll() ([]domain.Country, error)
}
