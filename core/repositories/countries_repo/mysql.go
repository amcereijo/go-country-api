package countries_repo

import (
	"fmt"

	"github.com/amcereijo/go-country-api/core/domain"
	"github.com/jinzhu/gorm"
)

type mysqlrepo struct {
	connector *gorm.DB
}

func New(connector *gorm.DB) *mysqlrepo {
	return &mysqlrepo{
		connector: connector,
	}
}

func (repo *mysqlrepo) SaveCountry(country domain.Country) (domain.Country, error) {
	fmt.Printf("To save country %+v\n", country)
	repo.connector.Create(&country)
	fmt.Printf("Saved country %+v\n", country)
	return country, nil
}

func (repo *mysqlrepo) GetAll() ([]domain.Country, error) {
	fmt.Println("Find all countries")
	var countriesEntity []domain.Country
	repo.connector.Find(&countriesEntity)
	fmt.Printf("Found %d countries\n", len(countriesEntity))

	return countriesEntity, nil
}
