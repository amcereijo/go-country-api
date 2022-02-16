package countries_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amcereijo/go-country-api/core/domain"
)

func Validator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		country := &domain.Country{}

		defer r.Body.Close()

		if err := json.NewDecoder(r.Body).Decode(country); err != nil {
			panic(err)
		}

		if validErrs := country.Validate(); len(validErrs) > 0 {
			err := map[string]interface{}{"validationError": validErrs}
			w.Header().Set("Content-type", "applciation/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			fmt.Printf("Returning validation errors: %v\n", err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
