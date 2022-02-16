package domain

import "net/url"

type Country struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
}

func (c *Country) Validate() url.Values {
	errs := url.Values{}

	if c.Name == "" {
		errs.Add("name", "invalid name")
	}
	if c.Capital == "" {
		errs.Add("capital", "invalid capital")
	}

	return errs
}
