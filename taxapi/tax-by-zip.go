package taxapi

import (
	"net/http"
	"time"
)

// SalesTaxByZip is the response object from /salestax/zip/{zip}
type SalesTaxByZip struct {
	Status string        `json:"status"`
	Rates  SalesTaxRates `json:"rates"`
}

// SalesTaxRates contains the tax rate data for a given zip
type SalesTaxRates struct {
	State                 string  `json:"state"`
	ZipCode               string  `json:"zipCode"`
	TaxRegionName         string  `json:"taxRegionName"`
	StateRate             float64 `json:"stateRate"`
	EstimatedCombinedRate float64 `json:"estimatedCombinedRate"`
	EstimatedCountyRate   float64 `json:"estimatedCountyRate"`
	EstimatedCityRate     float64 `json:"estimatedCityRate"`
	EstimatedSpecialRate  int64   `json:"estimatedSpecialRate"`
	RiskLevel             int64   `json:"riskLevel"`
}

// GetSalesTaxByZip will return sales tax data given a valid Zip Code
func (c *Client) GetSalesTaxByZip(zip string) (SalesTaxByZip, error) {
	cacheKey := "salesTax_" + zip
	result := new(SalesTaxByZip)

	// If caching is enabled, attempt to first get the result from cache
	if c.cache != nil {
		cacheResult, cacheHit := c.cache.Get(cacheKey)

		if cacheHit {
			result = cacheResult.(*SalesTaxByZip)
			return *result, nil
		}
	}

	_, err := c.makeRequest("/salestax/zip/"+zip, http.MethodGet, [][]string{}, result)

	// If caching is enabled, store the result in cache for a day
	if c.cache != nil {
		c.cache.Set(cacheKey, result, time.Hour*24)
	}

	return *result, err
}
