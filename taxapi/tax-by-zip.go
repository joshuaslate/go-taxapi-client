package taxapi

import "net/http"

// SalesTaxByZip is the response object from /salestax/zip/{zip}
type SalesTaxByZip struct {
	Status string `json:"status"`
	Rates  Rates  `json:"rates"`
}

// Rates contains the tax rate data for a given zip
type Rates struct {
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
	result := new(SalesTaxByZip)
	_, err := c.makeRequest("/salestax/zip/"+zip, http.MethodGet, [][]string{}, result)

	return *result, err
}
