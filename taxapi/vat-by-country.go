package taxapi

import (
	"net/http"
	"reflect"
)

// VATResponse is the response from /vat/rates
type VATResponse struct {
	Status string   `json:"status"`
	Rates  VATRates `json:"rates"`
}

// VATRates contains the rates for each country in the EU
type VATRates struct {
	AT CountryVATRates `json:"AT"`
	BE CountryVATRates `json:"BE"`
	BG CountryVATRates `json:"BG"`
	HR CountryVATRates `json:"HR"`
	CY CountryVATRates `json:"CY"`
	CZ CountryVATRates `json:"CZ"`
	DK CountryVATRates `json:"DK"`
	EE CountryVATRates `json:"EE"`
	FI CountryVATRates `json:"FI"`
	FR CountryVATRates `json:"FR"`
	DE CountryVATRates `json:"DE"`
	GR CountryVATRates `json:"GR"`
	HU CountryVATRates `json:"HU"`
	IE CountryVATRates `json:"IE"`
	IT CountryVATRates `json:"IT"`
	LV CountryVATRates `json:"LV"`
	LT CountryVATRates `json:"LT"`
	LU CountryVATRates `json:"LU"`
	MT CountryVATRates `json:"MT"`
	NL CountryVATRates `json:"NL"`
	PL CountryVATRates `json:"PL"`
	PT CountryVATRates `json:"PT"`
	RO CountryVATRates `json:"RO"`
	SK CountryVATRates `json:"SK"`
	SI CountryVATRates `json:"SI"`
	ES CountryVATRates `json:"ES"`
	SE CountryVATRates `json:"SE"`
	GB CountryVATRates `json:"GB"`
}

// CountryVATRates are the VAT rates for a single country
type CountryVATRates struct {
	CountryName  string       `json:"country_name"`
	StandardRate float64      `json:"standard_rate"`
	ReducedRates ReducedRates `json:"reduced_rates"`
}

// ReducedRates are rate reductions that may apply over the standard VAT rate
type ReducedRates struct {
	Accommodation                  *float64 `json:"accommodation"`
	AdmissionToCulturalEvents      *float64 `json:"admission_to_cultural_events"`
	AdmissionToEntertainmentEvents *float64 `json:"admission_to_entertainment_events"`
	AdmissionToSportingEvents      *float64 `json:"admission_to_sporting_events"`
	Advertising                    *float64 `json:"advertising"`
	AgriculturalSupplies           *float64 `json:"agricultural_supplies"`
	BabyFoodstuffs                 *float64 `json:"baby_foodstuffs"`
	Bikes                          *float64 `json:"bikes"`
	Books                          *float64 `json:"books"`
	ChildrensClothing              *float64 `json:"childrens_clothing"`
	DomesticFuel                   *float64 `json:"domestic_fuel"`
	DomesticServices               *float64 `json:"domestic_services"`
	EBooks                         *float64 `json:"e-books"`
	Foodstuffs                     *float64 `json:"foodstuffs"`
	Hotels                         *float64 `json:"hotels"`
	Medical                        *float64 `json:"medical"`
	Newspapers                     *float64 `json:"newspapers"`
	PassengerTransport             *float64 `json:"passenger_transport"`
	Pharmaceuticals                *float64 `json:"pharmaceuticals"`
	PropertyRenovations            *float64 `json:"property_renovations"`
	Restaurants                    *float64 `json:"restaurants"`
	SocialHousing                  *float64 `json:"social_housing"`
	Water                          *float64 `json:"water"`
	Wine                           *float64 `json:"wine"`
}

// GetAllVATRates returns VAT rates for all participating EU countries
func (c *Client) GetAllVATRates() (VATResponse, error) {
	result := new(VATResponse)
	_, err := c.makeRequest("/vat/rates", http.MethodGet, [][]string{}, result)

	return *result, err
}

// GetVATByCountryCode returns VAT rates by a given country code (DE, AT, ES, etc.)
func (c *Client) GetVATByCountryCode(countryCode string) (CountryVATRates, error) {
	vatResponse, err := c.GetAllVATRates()

	if err != nil {
		return CountryVATRates{}, err
	}

	reflection := reflect.ValueOf(vatResponse.Rates)
	country := reflect.Indirect(reflection).FieldByName(countryCode)

	return country.Interface().(CountryVATRates), nil
}
