package countries_service_test

import (
	"testing"

	"github.com/amcereijo/go-country-api/core/domain"
	"github.com/amcereijo/go-country-api/core/services/countries_service"
	mock_ports "github.com/amcereijo/go-country-api/mocks"
	"github.com/golang/mock/gomock"
)

func TestGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repomock := mock_ports.NewMockCountriesRepository(mockCtrl)
	repomock.EXPECT().
		GetAll().
		Return([]domain.Country{
			{
				ID:      1,
				Name:    "nombre",
				Capital: "capital",
			},
			{
				ID:      2,
				Name:    "nombre2",
				Capital: "capital2",
			},
		}, nil)

	service := countries_service.New(repomock)
	countries, _ := service.GetAll()

	if countries[0].ID != 1 || countries[0].Name != "nombre" || countries[0].Capital != "capital" ||
		countries[1].ID != 2 || countries[1].Name != "nombre2" || countries[1].Capital != "capital2" {
		t.Errorf("Error getting countries%v", countries)
	}
}

func TestCreateCountry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repomock := mock_ports.NewMockCountriesRepository(mockCtrl)
	repomock.
		EXPECT().
		SaveCountry(gomock.AssignableToTypeOf(domain.Country{})).
		DoAndReturn((func(c domain.Country) (domain.Country, error) {
			c.ID = 1
			return c, nil
		}))

	service := countries_service.New(repomock)
	country, _ := service.Create("Name", "Capital")

	if country.ID != 1 || country.Capital != "Capital" || country.Name != "Name" {
		t.Errorf("Error saving country %v", country)
	}
}
