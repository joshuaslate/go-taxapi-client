package taxapi

import (
	"net/http"
	"time"
)

// VATVerificationResponse is the response from vat?vat_number={vatNumber}
type VATVerificationResponse struct {
	Status         string `json:"status"`
	Valid          bool   `json:"valid"`
	FormatValid    bool   `json:"format_valid"`
	Query          bool   `json:"query"`
	CountryCode    string `json:"country_code"`
	VatNumber      string `json:"vat_number"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
}

// VerifyVATNumber will check that a VAT number is valid
func (c *Client) VerifyVATNumber(vatNumber string) (VATVerificationResponse, error) {
	cacheKey := "vatVerification_" + vatNumber
	vatQuery := []string{"vat_number", vatNumber}
	result := new(VATVerificationResponse)

	// If caching is enabled, attempt to first get the result from cache
	if c.cache != nil {
		cacheResult, cacheHit := c.cache.Get(cacheKey)

		if cacheHit {
			result = cacheResult.(*VATVerificationResponse)
			return *result, nil
		}
	}

	// If caching is enabled, store the result in cache for a day
	if c.cache != nil {
		c.cache.Set(cacheKey, result, time.Hour*24)
	}

	_, err := c.makeRequest("/vat", http.MethodGet, [][]string{vatQuery}, result)

	return *result, err
}
