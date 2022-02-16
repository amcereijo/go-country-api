package countries_service

import (
	"errors"

	"github.com/amcereijo/go-country-api/core/domain"
	"github.com/amcereijo/go-country-api/core/ports"
)

type service struct {
	countriesRepository ports.CountriesRepository
}

func New(countriesRepository ports.CountriesRepository) *service {
	return &service{
		countriesRepository: countriesRepository,
	}
}

func (srv *service) GetAll() ([]domain.Country, error) {
	countries, error := srv.countriesRepository.GetAll()
	if error != nil {
		return nil, error
	}
	return countries, nil
}

func (srv *service) Create(name string, capital string) (domain.Country, error) {
	if name == "" {
		return domain.Country{}, errors.New(("invalid name"))
	}
	if capital == "" {
		return domain.Country{}, errors.New(("invalid capital"))
	}
	country := domain.Country{Name: name, Capital: capital}
	var err error

	if country, err = srv.countriesRepository.SaveCountry(country); err != nil {
		return domain.Country{}, errors.New("error creating Country")
	}

	return country, nil
}
